package scripts

import (
	"context"
	"github.com/rarimo/near-go/common"
	nearclient2 "github.com/rarimo/near-go/nearclient"
	"github.com/rarimo/near-go/nearclient/models"
)

func FtDeposit(ctx context.Context, cli nearclient2.Client, sender, receiver, token string, amount string, bridge string, isWrapped bool) (string, string) {
	amn := parseAmount(amount)

	transferArgs := common.NewTransferArgs(token, sender, receiver, targetNetwork, isWrapped)
	msg, err := transferArgs.String()
	if err != nil {
		panic(err)
	}

	depositResp, err := cli.TransactionSendAwait(ctx, sender, token, []common.Action{
		common.NewFtTransferCall(common.FtTransferArgs{
			ReceiverId: bridge,
			Amount:     amn,
			Msg:        msg,
		}, MaxGas),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	eventID, err := GetDepositedReceiptID(depositResp, models.LogEventTypeFtDeposited, bridge, &token, nil, &amn)
	if err != nil {
		panic(err)
	}
	if eventID == nil {
		panic("eventID is nil")
	}

	return depositResp.Transaction.Hash.String(), eventID.String()
}
