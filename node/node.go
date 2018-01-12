package node

import (
	"github.com/ofbank_wallet/OFBANK_WALLET/rpc"
	"sync"
	"fmt"
	"github.com/ofbank_wallet/OFBANK_WALLET/log"
	"os"
	"github.com/ofbank_wallet/OFBANK_WALLET/accounts/keystore"
	"github.com/ofbank_wallet/OFBANK_WALLET/accounts"
)

type Node struct {

	lock sync.RWMutex

	inprocHandler *rpc.Server
    am *accounts.Manager
	stop chan struct{} // Channel to wait for termination notifications
}

var node *Node =nil

func NewNode() (*Node,error){

	if node==nil {
		newNode := new(Node)
		node  = newNode
		am,err:=makeAccountManager()
		if err!=nil{
         return  nil,err
		}
		node.am =am
		newNode.StartInProc(newNode.Apis())
	}
	return node,nil
}

func (n *Node) Attach() (*rpc.Client, error) {

	n.lock.RLock()
	defer n.lock.RUnlock()

	return rpc.DialInProc(n.inprocHandler), nil
}



func (n *Node) StartInProc(apis []rpc.API) error {
	// Register all the APIs exposed by the services
	handler := rpc.NewServer()
	for _, api := range apis {
		if err := handler.RegisterName(api.Namespace, api.Service); err != nil {
			return err
		}
		log.Debug(fmt.Sprintf("InProc registered %T under '%s'", api.Service, api.Namespace))
	}
	n.inprocHandler = handler
	return nil
}


func (n *Node) Apis() []rpc.API {

	return []rpc.API{
		{
			Namespace: "ofbank",
			Version:   "1.0",
			Service:   NewOFBankAPI(n.am),
		},
	}
}

func (n *Node) Wait() {
	n.lock.RLock()

	stop := n.stop
	n.lock.RUnlock()

	<-stop
}
func makeAccountManager() (*accounts.Manager, error) {

	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP

	var (
		keydir    string
		err       error
	)
	keydir = "./keystore"

	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(keydir, 0700); err != nil {
		return nil, err
	}
	// Assemble the account manager and supported backends
	backends := []accounts.Backend{
		keystore.NewKeyStore(keydir, scryptN, scryptP),
	}

	return accounts.NewManager(backends...), nil
}
