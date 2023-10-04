package constants

import "github.com/rarimo/near-go/common"

var (
	// 30 TGas
	DefaultFunctionCallGas common.Gas = 30 * 1000000000000
)

var (
	OrdMappings = map[string]uint8{
		"CreateAccount":  common.OrdCreateAccount,
		"DeployContract": common.OrdDeployContract,
		"FunctionCall":   common.OrdFunctionCall,
		"Transfer":       common.OrdTransfer,
		"Stake":          common.OrdStake,
		"AddKey":         common.OrdAddKey,
		"DeleteKey":      common.OrdDeleteKey,
		"DeleteAccount":  common.OrdDeleteAccount,
	}

	SimpleActions = map[string]bool{
		"CreateAccount": true,
	}
)
