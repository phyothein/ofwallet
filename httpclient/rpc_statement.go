package httpclient

import (
	"github.com/ofbank_wallet/OFBANK_WALLET/common"
	"github.com/ofbank_wallet/OFBANK_WALLET/common/hexutil"

	"math/big"
	"encoding/json"
	"github.com/ofbank_wallet/OFBANK_WALLET"
	"strconv"
)


type RPCStringResult struct {
	JsonRPC string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
	Errors  *Error `    json:"error"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func CreateGETNonceState(address common.Address) string {

	addressString := hexutil.Encode(address[:])

	return "{\"id\":1,\"method\":\"ofbank_getNonce\",\"params\":[" + "\"" + addressString + "\"" + "]}"

}

func CreateGetBalanceStatement(address common.Address) string {

	addressString := hexutil.Encode(address[:])

	return "{\"id\":1,\"method\":\"ofbank_show\",\"params\":[" + "\"" + addressString + "\"," + "\"latest\"]}"

}
func CreateGetBlockStatement(number int) string {

	return "{\"id\":1,\"method\":\"ofbank_getBlock\",\"params\":[" + strconv.Itoa(number) + "]}"

}

func CreateTransacationByHash(hash string) string {

	return "{\"id\":1,\"method\":\"ofbank_checkTrans\",\"params\":[" + "\"" + hash + "\"" + "]}"

}
func CreateGetBlockHeight() string {

	return "{\"id\":1,\"method\":\"ofbank_lastBN\",\"params\":[]}"

}

func CreateTransacationsByAddressStatement(address common.Address) string {

	addressString := hexutil.Encode(address[:])

	return "{\"id\":1,\"method\":\"ofbank_checkTransFrom\",\"params\":[" + "\"" + addressString + "\"" + "]}"
}

type TXJson struct {
	AccountNonce uint64 `json:"nonce"    gencodec:"required"`
	Price        string `json:"gasPrice" gencodec:"required"`
	GasLimit     string `json:"gas"      gencodec:"required"`
	Recipient    string `json:"to"       rlp:"nil"` // nil means contract creation
	from         string `json:"from"       rlp:"required"`
	Amount       string `json:"value"    gencodec:"required"` // Water Egg
	V            string `json:"v" gencodec:"required"`
	R            string `json:"r" gencodec:"required"`
	S            string `json:"s" gencodec:"required"`
	ACcode       string `json:"c" gencodec:"required"`
}

func CreatePushTxStatement(acountNonce uint64, price, gasLimit, amount *big.Int, to, from *common.Address, v, s, r *big.Int, accode []byte) string {

	priceString := price.String()
	gasLimitString := gasLimit.String()
	amountString := amount.String()
	toString := hexutil.Encode(to[:])
	fromString := hexutil.Encode(from[:])
	vString := v.String()
	sString := s.String()
	rString := r.String()
	accodeString := hexutil.Encode(accode)

	txJson := &TXJson{
		AccountNonce: acountNonce,
		Price:        priceString,
		GasLimit:     gasLimitString,
		Recipient:    toString,
		from:         fromString,
		Amount:       amountString,
		V:            vString,
		R:            rString,
		S:            sString,
		ACcode:       accodeString,
	}

	jsonString, err := json.Marshal(txJson)
	if err != nil {
		ofbank_wallet.Logger.Error("Fail to marshal Json")
		return "'"
	}

	return "{\"id\":1,\"method\":\"ofbank_pushTX\",\"params\":[" + string(jsonString) + "]}"

}
