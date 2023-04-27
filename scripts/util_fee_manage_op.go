package scripts

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/mr-tron/base58"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
	"math/big"
	"time"
)

func getFeeManageOperationSignArgs(operationType action.FeeManageOperationType, token action.FeeToken, feeAmount, privateKey, receiver, bridgeAddr string) (
	originHash string,
	signature string,
	resultPath [][32]byte,
	recoveryID byte,
) {
	content := []byte{byte(operationType)}

	if token.TokenType != action.TokenType_Native && token.TokenAddr != nil {
		content = append(content, hexutil.MustDecode(hexutil.Encode([]byte(*token.TokenAddr)))...)
	}

	content = append(content, to32Bytes(amountBytes(feeAmount))...)

	fmt.Println("operation data", base58.Encode(content))

	return getContent(
		privateKey,
		time.Now().String(),
		"",
		receiver,
		bridgeAddr,
		"Near",
		"Near",
		content,
	)
}

func amountBytes(amount string) []byte {
	am, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return []byte{}
	}

	return am.Bytes()
}
