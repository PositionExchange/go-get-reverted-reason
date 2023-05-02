# Go Get Revert Reason
This package provides a simple way to retrieve the revert reason for a failed Ethereum transaction. It can be easily integrated into your Go projects to interact with the Ethereum blockchain.

# Installation
To install the get-revert-reason package, run the following command:

```
go get github.com/PositionExchange/go-get-reverted-reason
```

# Usage
Here's a basic example of how to use the get-revert-reason package:

```go
package main

import (
	"context"
	"fmt"

	"github.com/PositionExchange/go-get-reverted-reason/revertreason"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	client, err := ethclient.Dial("RPC_URL")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Replace with a valid transaction hash
	txHash := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		panic(err)
	}

	revertReason, err := revertreason.GetRevertReason(client, tx, context.Background())
	if err != nil {
		panic(err)
	}

	if revertReason != "" {
		fmt.Printf("Revert reason: %s\n", revertReason)
	} else {
		fmt.Println("Transaction did not revert or revert reason not found")
	}
}
```
Remember to replace RPC_URL with the actual URL of your Ethereum JSON-RPC provider.

# Contributing
Contributions to the get-revert-reason package are welcome. Please submit a pull request or open an issue on GitHub to contribute.

# License
This package is released under the MIT License.
