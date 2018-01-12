package types

import (
	"github.com/ofbank_wallet/OFBANK_WALLET/common"
	"github.com/ofbank_wallet/OFBANK_WALLET/crypto/sha3"
	"github.com/ofbank_wallet/OFBANK_WALLET/rlp"
)
type writeCounter common.StorageSize

func (c *writeCounter) Write(b []byte) (int, error) {
	*c += writeCounter(len(b))
	return len(b), nil
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}
