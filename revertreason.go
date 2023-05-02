package revertreason

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetRevertReason(client *ethclient.Client, hash common.Hash) (string, error) {
    tx, _, err := client.TransactionByHash(context.Background(), hash)
    if err != nil {
        return "", err
    }

    from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
    if err != nil {

			return "", err
		}


    msg := ethereum.CallMsg{
        From:     from,
        To:       tx.To(),
        Gas:      tx.Gas(),
        GasPrice: tx.GasPrice(),
        Value:    tx.Value(),
        Data:     tx.Data(),
    }

    res, err := client.CallContract(context.Background(), msg, nil)
    if err != nil {
				if strings.Contains(err.Error(), "reverted"){
						return err.Error(), nil
				}	
        return "", err
    }

    return string(res), nil
}


func parseRevertReason(input []byte) (string, error) {
	if len(input) < 4 {
		return "", fmt.Errorf("invalid input")
	}

	// methodID := input[:4]
	inputData := input[4:]

	parsedRevert, err := abi.JSON(strings.NewReader(abiRevert))
	if err != nil {
		return "", err
	}

	var reason string
	err = parsedRevert.UnpackIntoInterface(&reason, "Error", inputData)
	if err != nil {
		return "", err
	}

	return reason, nil
}

const abiRevert = `[{ "name": "Error", "type": "function", "inputs": [ { "type": "string" } ] }]`

