package common

var (
	NftMintStorageDeposit = MustBalanceFromString("0.2")
	FtMintStorageDeposit  = MustBalanceFromString("0.00125")
)

const (
	ContractNew = "new"

	ContractNftTransfer       = "nft_transfer_call"
	ContractNftWithdraw       = "nft_withdraw"
	ContractNftTokensForOwner = "nft_tokens_for_owner"
	ContractNftMetadata       = "nft_metadata"
	ContractNftGet            = "nft_token"

	ContractFtTransfer  = "ft_transfer_call"
	ContractFtWithdraw  = "ft_withdraw"
	ContractFtBalanceOf = "ft_balance_of"
	ContractFtMint      = "ft_mint"

	ContractStorageDeposit = "storage_deposit"

	ContractBridgeNativeDeposit  = "native_deposit"
	ContractBridgeNativeWithdraw = "native_withdraw"

	ContractFeerRegister       = "register"
	ContractFeerUnregister     = "unregister"
	ContractFeerChargeNative   = "charge_native"
	ContractFeerAddFeeToken    = "add_fee_token"
	ContractFeerUpdateFeeToken = "update_fee_token"
	ContractFeerRemoveFeeToken = "remove_fee_token"
	ContractFeerWithdraw       = "withdraw"
	ContractFeerGetFeeToken    = "get_fee_token"
	ContractFeerGetFeeTokens   = "get_fee_tokens"
)
