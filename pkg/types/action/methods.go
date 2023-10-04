package action

const (
	NewMethod = "new"

	NftTransferMethod         = "nft_transfer_call"
	NftWithdrawMethod         = "nft_withdraw"
	NftTokensForOwnerMethod   = "nft_tokens_for_owner"
	NftContractMetadataMethod = "nft_metadata"
	NftGetMethod              = "nft_token"

	FtTransferMethod       = "ft_transfer_call"
	FtWithdrawMethod       = "ft_withdraw"
	FtBalanceOfMethod      = "ft_balance_of"
	FtStorageDepositMethod = "storage_deposit"
	FtMintMethod           = "ft_mint"

	BridgeNativeDepositMethod  = "native_deposit"
	BridgeNativeWithdrawMethod = "native_withdraw"

	FeerRegister       = "register"
	FeerUnregister     = "unregister"
	FeerChargeNative   = "charge_native"
	FeerAddFeeToken    = "add_fee_token"
	FeerUpdateFeeToken = "update_fee_token"
	FeerRemoveFeeToken = "remove_fee_token"
	FeerWithdraw       = "withdraw"
	FeerGetFeeToken    = "get_fee_token"
	FeerGetFeeTokens   = "get_fee_tokens"
)
