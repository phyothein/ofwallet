package node

import (
	"fmt"
	"github.com/ofbank_wallet/OFBANK_WALLET/common"
	"regexp"
	"strconv"
	"github.com/ofbank_wallet/OFBANK_WALLET/accounts"
	"github.com/ofbank_wallet/OFBANK_WALLET/accounts/keystore"
	"math/big"
	"github.com/ofbank_wallet/OFBANK_WALLET/common/hexutil"
	"context"
	"github.com/ofbank_wallet/OFBANK_WALLET/httpclient"
	"encoding/json"
	"github.com/ofbank_wallet/OFBANK_WALLET/core/types"
	"time"
	"errors"
	"github.com/ofbank_wallet/OFBANK_WALLET/common/math"
)

const (
	DEFAULTGAS  = 50000
	SUPERNODEIP = "http://127.0.0.1:8888"
)

type OFBankAPI struct {
	am *accounts.Manager
}

func NewOFBankAPI(am *accounts.Manager) *OFBankAPI {
	return &OFBankAPI{am: am}
}

//------------------------------------------------API 注册接口--------------------------------------------------
func (api *OFBankAPI) Register(sccode string, password string) (common.Address, error) {

	if !isNumber(sccode) {
		return common.Address{}, fmt.Errorf("contry code is not number")
	}

	i64_ccode, _ := strconv.ParseInt(sccode, 10, 64)

	ccode := int(i64_ccode)

	acode := int(6)

	cbt := []int{0, 0, 0, 0, 0} // get reginfo from central db
	if ccode > 0 && ccode < 1024 {
		cbt[0] = acode / 256 / 256
		cbt[1] = (acode / 256) % 256
		cbt[2] = acode % 256
		cbt[3] = ccode / 256
		cbt[4] = ccode % 256
	} else {
		cbt = []int{0, 0, 0, 0, 156}
	}

	acc, err := fetchKeystore(api.am).NewAccount(cbt, password)

	if err != nil {
		return common.Address{}, err
	}
	return acc.Address, nil
}


//------------------------------------------------API 发送交易接口--------------------------------------------------
func (api *OFBankAPI) SendTranToNode(ctx context.Context, trans SendTxArgs1) (string, error) {

	acc_f := accounts.Account{Address: trans.From}

	transaction, err := aTa1(trans)
	if err != nil {
		return "", err
	}

	wallet, err := api.am.Find(acc_f)

	if err != nil {
		return "", err
	}

	if transaction.To != nil {
		if transaction.From == *transaction.To {
			return "", nil
		}
	}

	if err := transaction.setDefaults(); err != nil {
		return "", err
	}

	tx := transaction.toTransaction()

	signed, err := wallet.SignTx(acc_f, tx, nil)

	if err != nil {
		return "", err
	}

	httpResq := httpclient.NewHttpRequest(SUPERNODEIP, httpclient.POST, httpclient.CreatePushTxStatement(signed.Nonce(),
		signed.GasPrice(),
		signed.Gas(),
		signed.Value(),
		signed.To(),
		&transaction.From,
		signed.GetV(),
		signed.GetS(),
		signed.GetR(),
		signed.GetAcc()))

	httpResq.Do()

	rep := <-httpResq.Response

	if rep.StatusCode != 200 {

		return "", &httpclient.FailtoConnectRPC{Message: "Fail to connect with suhper nodes...."}
	}

	StringResult := new(RPCStringResult)

	err = json.Unmarshal([]byte(rep.Body), StringResult)
	if err != nil {
		return "", err
	}

	if StringResult.Error != nil {
		return "", &httpclient.FailtoConnectRPC{Message: StringResult.Error.Message}
	}

	return StringResult.Result, nil
}

func aTa1(args SendTxArgs1) (SendTxArgs, error) {
	var arg SendTxArgs

	arg.Gas = args.Gas
	arg.useTime = args.UseTime
	arg.To = args.To
	arg.Data = args.Data
	arg.GasPrice = args.GasPrice
	arg.From = args.From
	arg.Nonce = args.Nonce

	m, err := strconv.ParseFloat(args.Value, 64)
	if err != nil {
		return arg, err
	}

	amount := new(big.Int).Mul(big.NewInt(int64(m*1E3)), big.NewInt(1e15))
	arg.Value = (*hexutil.Big)(amount)

	return arg, nil
}

func (args *SendTxArgs) toTransaction() *types.Transaction {
	abt := args.From[:5]
	var cbt = []byte{0, 0, 0, 0, 156}
	cbt[0] = byte(abt[0])
	cbt[1] = byte(abt[1])
	cbt[2] = byte(abt[2])
	cbt[3] = byte(abt[3])
	cbt[4] = byte(abt[4])
	if args.To == nil {
		return types.NewContractCreation(uint64(*args.Nonce), (*big.Int)(args.Value), (*big.Int)(args.Gas), (*big.Int)(args.GasPrice), (*big.Int)(args.useTime), args.Data, cbt)
	}
	return types.NewTransaction(uint64(*args.Nonce), *args.To, (*big.Int)(args.Value), (*big.Int)(args.Gas), (*big.Int)(args.GasPrice), (*big.Int)(args.useTime), args.Data, cbt)
}

// prepareSendTxArgs is a helper function that fills in default values for unspecified tx fields.
func (args *SendTxArgs) setDefaults() error {
	if args.Gas == nil {
		args.Gas = (*hexutil.Big)(big.NewInt(DEFAULTGAS))
	}
	//可以发送请求
	if args.GasPrice == nil {
		price := new(big.Int).SetInt64(190)

		args.GasPrice = (*hexutil.Big)(price)
	}
	if args.Value == nil {
		args.Value = new(hexutil.Big) // Water Egg
	}

	if args.Nonce == nil {
		nonce, err := getNonce(args.From)

		if err != nil {
			return err
		}
		args.Nonce = (*hexutil.Uint64)(&nonce)
	}
	return nil
}

func getNonce(address common.Address) (uint64, error) {

	httpResq := httpclient.NewHttpRequest(SUPERNODEIP, httpclient.POST, httpclient.CreateGETNonceState(address))

	httpResq.Do()

	rep := <-httpResq.Response

	if rep.StatusCode != 200 {
		return 0, &httpclient.FailtoConnectRPC{Message: "Fail to connect with suhper nodes...."}
	}
	var rpcResut RPCIntReslt
	err := json.Unmarshal([]byte(rep.Body), &rpcResut)

	if err != nil {
		return 0, &httpclient.FailtoConnectRPC{Message: rpcResut.Error.Message}
	}

	if rpcResut.Error!= nil {
		return 0, &httpclient.FailtoConnectRPC{Message: rpcResut.Error.Message}
	}

	return rpcResut.Result, nil
}

//------------------------------------------------API 解锁接口--------------------------------------------------
func (s *OFBankAPI) UnlockW(addr common.Address, password string, duration *uint64) (bool, error) {

	const max = uint64(time.Duration(math.MaxInt64) / time.Second)
	var d time.Duration
	if duration == nil {
		d = 120 * time.Second
	} else if *duration > max {
		return false, errors.New("unlock duration too large")
	} else {
		d = time.Duration(*duration) * time.Second
	}
	err := fetchKeystore(s.am).TimedUnlock(accounts.Account{Address: addr}, password, d)
	return err == nil, err
}


//------------------------------------------------API 多次签名--------------------------------------------------
func (api *OFBankAPI) RegisterWithMulKeys(sccode string, count int, password string) ([]common.Address, error) {
	keys := make([]common.Address, 0)

	if !isNumber(sccode) {
		return keys, fmt.Errorf("contry code is not number")
	}


	for i := 0; i < count; i++ {
		key, err := api.Register(sccode, password)
		if err != nil {
			continue
		}
		fmt.Println(hexutil.Encode(key[:]))

		keys = append(keys, key)

		}

	return keys, nil
}


//------------------------------------------------API 通过高度获取区块--------------------------------------------------
func (api *OFBankAPI) GetBlock(number int) (*RPCBlock, error) {

	httpResq := httpclient.NewHttpRequest(SUPERNODEIP, httpclient.POST, httpclient.CreateGetBlockStatement(number))

	httpResq.Do()

	rep := <-httpResq.Response

	if rep.StatusCode != 200 {
		return nil, &httpclient.FailtoConnectRPC{Message: "Fail to connect with suhper nodes...."}
	}

	block := new(RPCTBlockResult)

	err := json.Unmarshal([]byte(rep.Body), block)

	if err != nil {
		return nil, err
	}

	if block.Error!=nil{
		return nil,&httpclient.FailtoConnectRPC{Message:block.Error.Message}
	}

   block.Result.BlokReward="913.242"

	return &block.Result, nil

}

//------------------------------------------------API 显示所有钱包地址--------------------------------------------------
func (s *OFBankAPI) Accounts() []common.Address {
	addresses := make([]common.Address, 0) // return [] instead of nil if empty
	for _, wallet := range s.am.Wallets() {
		for _, account := range wallet.Accounts() {
			addresses = append(addresses, account.Address)
		}
	}
	return addresses
}



//------------------------------------------------API 通过交易Hash获取交易--------------------------------------------------
func (api *OFBankAPI) GetTransactionByHash(hash string) (*RPCStringTransaction, error) {

	httpResq := httpclient.NewHttpRequest(SUPERNODEIP, httpclient.POST, httpclient.CreateTransacationByHash(hash))

	httpResq.Do()

	rep := <-httpResq.Response

	rpcTran := new(RPCTransactionResult)

	if rep.StatusCode != 200 {
		return nil, &httpclient.FailtoConnectRPC{Message: "Fail to connect with suhper nodes...."}
	}
	err := json.Unmarshal([]byte(rep.Body), rpcTran)

	if err != nil {
		return nil, err
	}

	if rpcTran.Error!=nil{
	   return nil,&httpclient.FailtoConnectRPC{Message:rpcTran.Error.Message}
	}

	return &rpcTran.Result, nil
}


//------------------------------------------------API 通过地址获取所有交易--------------------------------------------------
func (api *OFBankAPI) GetTransactionsByAddress(address common.Address) ([]RPCStringTransaction, error) {

	transcations := new(RPCTranscations)

	httpResq := httpclient.NewHttpRequest(SUPERNODEIP, httpclient.POST, httpclient.CreateTransacationsByAddressStatement(address))

	httpResq.Do()

	rep := <-httpResq.Response

	if rep.StatusCode != 200 {
		return nil, &httpclient.FailtoConnectRPC{Message: "Fail to connect with suhper nodes...."}
	}

	err := json.Unmarshal([]byte(rep.Body), transcations)

	if err != nil {
		return nil, &httpclient.FailtoConnectRPC{Message: "Fail to Unmarshal to Json...."}
	}

	if transcations.Error!=nil{
		return  nil,&httpclient.FailtoConnectRPC{Message:transcations.Error.Message}
	}




	return transcations.Result, nil

}

//------------------------------------------------API 获取当前区块高度--------------------------------------------------
func (api *OFBankAPI) GetBlockHeight()(string,error){


	httpResq := httpclient.NewHttpRequest(SUPERNODEIP, httpclient.POST, httpclient.CreateGetBlockHeight())

	httpResq.Do()

	rep := <-httpResq.Response


	if rep.StatusCode != 200 {
		return  "", &httpclient.FailtoConnectRPC{Message: "Fail to connect with suhper nodes...."}
	}

	heigth :=new(RPCStringResult)

	err:=json.Unmarshal([]byte(rep.Body),heigth)

	if err!=nil{
		return  "",err

	}

	if heigth.Error!=nil{
		return "", &httpclient.FailtoConnectRPC{Message:heigth.Error.Message}
	}

	heightInt,_:=strconv.ParseInt(heigth.Result[2:],16,64)

	return strconv.Itoa(int(heightInt)),nil
}


//------------------------------------------------API 指定账户余额--------------------------------------------------
func (api *OFBankAPI) GetBalance(address common.Address) (string, error) {

	httpResq := httpclient.NewHttpRequest(SUPERNODEIP, httpclient.POST, httpclient.CreateGetBalanceStatement(address))

	httpResq.Do()

	rep := <-httpResq.Response

	if rep.StatusCode != 200 {
		return "", &httpclient.FailtoConnectRPC{Message: "Fail to connect with suhper nodes...."}
	}

	var rpcResut httpclient.RPCStringResult
	err := json.Unmarshal([]byte(rep.Body), &rpcResut)

	if err != nil {
		return "", &httpclient.FailtoConnectRPC{Message: "Wrong Type"}
	}

	if rpcResut.Errors != nil {

		return "", &httpclient.FailtoConnectRPC{Message: rpcResut.Errors.Message}
	}

	return rpcResut.Result, nil
}

// fetchKeystore retrives the encrypted keystore from the account manager.
func fetchKeystore(am *accounts.Manager) *keystore.KeyStore {
	return am.Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
}

//nancy
func isNumber(appcode string) bool {
	if appcode == "" {
		return false
	}
	str := fmt.Sprintf("^[0-9]{%d}$", len(appcode))
	m, err := regexp.MatchString(str, appcode)
	if err != nil {
		return false
	}
	return m
}
