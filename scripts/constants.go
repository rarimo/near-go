package scripts

import (
	"crypto/rand"
	"fmt"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"

	xcrypto "gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/origin"
)

var network = types.NetworkTestnet

var (
	content1 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash1", "networkFrom1", "eventId1").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  network,
		TargetContract: getRandSlice(),
		Data: operation.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "").GetContent(),
	}

	content2 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash2", "networkFrom2", "eventId2").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  network,
		TargetContract: getRandSlice(),
		Data: operation.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "").GetContent(),
	}

	content3 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash3", "networkFrom3", "eventId3").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  network,
		TargetContract: getRandSlice(),
		Data: operation.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "").GetContent(),
	}

	content4 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash3", "networkFrom3", "eventId3").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  network,
		TargetContract: getRandSlice(),
		Data: operation.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "").GetContent(),
	}

	content5 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash5", "networkFrom5", "eventId5").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  network,
		TargetContract: getRandSlice(),
		Data: operation.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "").GetContent(),
	}

	content6 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash6", "networkFrom6", "eventId6").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  network,
		TargetContract: getRandSlice(),
		Data: operation.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "").GetContent(),
	}

	content7 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash7", "networkFrom7", "eventId7").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  network,
		TargetContract: getRandSlice(),
		Data: operation.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "").GetContent(),
	}

	content8 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash8", "networkFrom8", "eventId8").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  network,
		TargetContract: getRandSlice(),
		Data: operation.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "").GetContent(),
	}

	content9 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash9", "networkFrom9", "eventId9").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  network,
		TargetContract: getRandSlice(),
		Data: operation.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "").GetContent(),
	}
)

func getRandSlice() []byte {
	var hash [32]byte
	rand.Read(hash[:])
	return hash[:]
}
