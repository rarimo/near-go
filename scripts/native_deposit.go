package scripts

//func NativeDeposit(ctx context.Context, cli client.Client, sender, receiver string, amount types.Amount, bridge string) (string, string) {
//	depositResp, err := cli.TransactionSendAwait(ctx, sender, bridge, []base.Action{
//		//action.NewNativeDepositCall(action.NativeDepositArgs{
//		//	ReceiverId: receiver,
//		//	Amount:     amount,
//		//}, MaxGas, amount),
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	eventID, err := GetDepositedReceiptID(depositResp, client.LogEventTypeNativeDeposited, bridge, bridge, nil, &amount)
//	if err != nil {
//		panic(err)
//	}
//	if eventID == nil {
//		panic("eventID is nil")
//	}
//	return depositResp.Transaction.Hash.String(), eventID.String()
//}
