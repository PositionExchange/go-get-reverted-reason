package revertreason_test

import (
	"testing"

	revertreason "github.com/PositionExchange/go-get-reverted-reason"
	"github.com/magiconair/properties/assert"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestGetRevertReasonPosiChain(t *testing.T) {
	// Replace with your Ethereum client URL
	client, err := ethclient.Dial("https://api.s0.t.posichain.org/")
	if err != nil {
		panic(err)
	}
	// Replace with the failed transaction hash
	txHash := common.HexToHash("0x3c2bacb88c4f5cb08fc88130d97b6b4ba978c49b11b12057e9f8357dd43e988e")
  revertReason, err := revertreason.GetRevertReason(client, txHash)
	t.Log(revertReason)
	assert.Equal(t, err, nil)
	assert.Equal(t, revertReason, "execution reverted: 26")
	// This transaction success
  revertReason, err = revertreason.GetRevertReason(client, common.HexToHash("0x3673103a600723d7122f6927bd179b181588073ea43b6fa979c464b8f15a6e71"))
	assert.Equal(t, err, nil)
	assert.Equal(t, revertReason, "")
}

func TestGetRevertReasonPosiChainBSC(t *testing.T) {
	// Replace with your Ethereum client URL
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		panic(err)
	}
	// Replace with the failed transaction hash
	txHash := common.HexToHash("0xdba10e26e2061bc3474b0d2ebcd666298da1bbc60b7adb07105dc022af6cfd2a")
  revertReason, err := revertreason.GetRevertReason(client, txHash)
	t.Log(revertReason)
	assert.Equal(t, err, nil)
	assert.Equal(t, revertReason, "execution reverted: 1")

	// This transaction success
  revertReason, err = revertreason.GetRevertReason(client, common.HexToHash("0x1eb9e64dca2c8bfd62a11f2a67172b2cbc4f4286f034dcd96f079918bc50d44f"))
	assert.Equal(t, err, nil)
	assert.Equal(t, revertReason, "")

}
