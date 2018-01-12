package node

import (
	"math/big"
	"github.com/ofbank_wallet/OFBANK_WALLET/common"
	"github.com/ofbank_wallet/OFBANK_WALLET/common/hexutil"
)

type SendTxArgs1 struct {
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to"`
	Gas      *hexutil.Big    `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    string 	     `json:"value"`				// Water Egg
	Data     hexutil.Bytes   `json:"data"`
	Nonce    *hexutil.Uint64 `json:"nonce"`
	UseTime  *hexutil.Big    `json:"useTime"`
}

type SendTxArgs struct {
	From     common.Address
	To       *common.Address
	Gas      *hexutil.Big
	GasPrice *hexutil.Big
	Value    *hexutil.Big
	Data     hexutil.Bytes
	Nonce    *hexutil.Uint64
	useTime  *hexutil.Big
}

type RPCBLockHeight struct {
	JsonRPC string                 `json:"jsonrpc"`
	Id      int                    `json:"id"`
	Result  []RPCStringTransaction `json:"result"`
	Error   ErrorMessage   `json:"error"`
}

type RPCTranscations struct {
	JsonRPC string                 `json:"jsonrpc"`
	Id      int                    `json:"id"`
	Result  []RPCStringTransaction `json:"result"`
	Error   *ErrorMessage   `json:"error"`
}

type RPCStringTransaction struct {
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	From             string   `json:"from"`
	Gas              string   `json:"gas"`
	GasPrice         string   `json:"gasPrice"`
	Hash             string   `json:"hash"`
	Input            string   `json:"input"`
	Nonce            string   `json:"nonce"`
	To               string   `json:"to"`
	TransactionIndex string   `json:"transactionIndex"`
	Value            *big.Int `json:"value"`
	V                string   `json:"v"`
	R                string   `json:"r"`
	S                string   `json:"s"`
}

type RPCTransaction struct {
	BlockHash        common.Hash     `json:"blockHash"`
	BlockNumber      *hexutil.Big    `json:"blockNumber"`
	From             common.Address  `json:"from"`
	Gas              *hexutil.Big    `json:"gas"`
	GasPrice         *hexutil.Big    `json:"gasPrice"`
	Hash             common.Hash     `json:"hash"`
	Input            hexutil.Bytes   `json:"input"`
	Nonce            hexutil.Uint64  `json:"nonce"`
	To               *common.Address `json:"to"`
	TransactionIndex hexutil.Uint    `json:"transactionIndex"`
	Value            *hexutil.Big    `json:"value"`
	V                *hexutil.Big    `json:"v"`
	R                *hexutil.Big    `json:"r"`
	S                *hexutil.Big    `json:"s"`
}

type RPCTransactionResult struct {
	JsonRPC string         `json:"jsonrpc"`
	Id      int            `json:"id"`
	Result  RPCStringTransaction `json:"result"`
	Error   *ErrorMessage   `json:"error"`
}
type RPCStringResult struct {
	JsonRPC string         `json:"jsonrpc"`
	Id      int            `json:"id"`
	Result  string         `json:"result"`
	Error   *ErrorMessage   `json:"error"`
}


type RPCIntReslt struct {
	JsonRPC string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  uint64 `json:"result"`
	Error   *ErrorMessage   `json:"error"`
}

type ErrorMessage struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type RPCTBlockResult struct {
	JsonRPC string   `json:"jsonrpc"`
	Id      int      `json:"id"`
	Result  RPCBlock `json:"result"`
	Error   *ErrorMessage   `json:"error"`

}

type RPCBlock struct {
	Coinbase         string            `json:"coinbase"`
	Difficult        string            `json:"difficulty"`
	ExtraData        string            `json:"extraData"`
	GasLimit         int               `json:"gasLimit"`
	GasUsed          string            `json:"gasUsed"`
	Hash             string            `json:"hash"`
	Miner            string            `json:"miner"`
	MixHash          string            `json:"mixHash"`
	Nonce            string            `json:"nonce"`
	Number           string            `json:"number"`
	ParentHash       string            `json:"parentHash"`
	ReceiptsRoot     string            `json:"receiptsRoot"`
	Sha3Uncles       string            `json:"sha3Uncles"`
	Size             string            `json:"size"`
	StateRoot        string            `json:"stateRoot"`
	Timestamp        string            `json:"timeStamp"`
	TotalDifficulty  string            `json:"totalDifficulty"`
	TransactionsRoot string            `json:"transactionsRoot"`
	BlokReward       string             `json:"blockReward"`
	Transactions     []TransactionHash `json:"transactions"`
}

type TransactionHash struct {
	Hash string `json:"hash"`
}
