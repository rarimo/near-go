package types

var (
	// 30 TGas
	DefaultFunctionCallGas Gas     = 30 * 1000000000000
	NftMintStorageDeposit  Balance = MustBalanceFromString("0.2")
	FtMintStorageDeposit   Balance = MustBalanceFromString("0.00125")
)
