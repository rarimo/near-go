package common

import (
	"encoding/json"
	"fmt"
	"github.com/near/borsh-go"
	"github.com/rarimo/near-go/constants"
	"lukechampine.com/uint128"
)

type Action struct {
	Enum borsh.Enum `borsh_enum:"true"`

	CreateAccount  ActionCreateAccount
	DeployContract ActionDeployContract
	FunctionCall   ActionFunctionCall
	Transfer       ActionTransfer
	Stake          ActionStake
	AddKey         ActionAddKey
	DeleteKey      ActionDeleteKey
	DeleteAccount  ActionDeleteAccount
}

const (
	OrdCreateAccount uint8 = iota
	OrdDeployContract
	OrdFunctionCall
	OrdTransfer
	OrdStake
	OrdAddKey
	OrdDeleteKey
	OrdDeleteAccount
)

func (a *Action) PrepaidGas() Gas {
	switch uint8(a.Enum) {
	case OrdFunctionCall:
		return a.FunctionCall.Gas
	default:
		return 0
	}
}

func (a *Action) DepositBalance() Balance {
	switch uint8(a.Enum) {
	case OrdFunctionCall:
		return a.FunctionCall.Deposit
	case OrdTransfer:
		return a.Transfer.Deposit
	default:
		return Balance(uint128.Zero)
	}
}

func (a *Action) UnderlyingValue() interface{} {
	switch uint8(a.Enum) {
	case OrdCreateAccount:
		return &a.CreateAccount
	case OrdDeployContract:
		return &a.DeployContract
	case OrdFunctionCall:
		return &a.FunctionCall
	case OrdTransfer:
		return &a.Transfer
	case OrdStake:
		return &a.Stake
	case OrdAddKey:
		return &a.AddKey
	case OrdDeleteKey:
		return &a.DeleteKey
	case OrdDeleteAccount:
		return &a.DeleteAccount
	}

	panic("unreachable")
}

func (a Action) String() string {
	ul := a.UnderlyingValue()
	if u, ok := ul.(interface{ String() string }); ok {
		return fmt.Sprintf("Action{%s}", u.String())
	}

	return fmt.Sprintf("Action{%#v}", ul)
}

func (a *Action) UnmarshalJSON(b []byte) (err error) {
	var obj map[string]json.RawMessage

	// actions can be either strings, or objects, so try deserializing into string first
	var actionType string
	if len(b) > 0 && b[0] == '"' {
		if err = json.Unmarshal(b, &actionType); err != nil {
			return
		}

		if _, ok := constants.SimpleActions[actionType]; !ok {
			err = fmt.Errorf("Action '%s' had no body", actionType)
			return
		}

		obj = map[string]json.RawMessage{
			actionType: json.RawMessage(`{}`),
		}
	} else {
		if err = json.Unmarshal(b, &obj); err != nil {
			return
		}
	}

	if l := len(obj); l > 1 {
		err = fmt.Errorf("action object contains invalid amount of keys (expected: 1, got: %d)", l)
		return
	}

	for k := range obj {
		actionType = k
		break
	}

	ord := constants.OrdMappings[actionType]
	*a = Action{Enum: borsh.Enum(ord)}
	ul := a.UnderlyingValue()

	if err = json.Unmarshal(obj[actionType], ul); err != nil {
		return
	}

	return nil
}

type ActionCreateAccount struct {
}

// NewCreateAccount Create an (sub)account using a transaction `receiver_id` as an ID for a new account
func NewCreateAccount() Action {
	return Action{
		Enum:          borsh.Enum(OrdCreateAccount),
		CreateAccount: ActionCreateAccount{},
	}
}

type ActionDeployContract struct {
	Code []byte `json:"code"`
}

func NewDeployContract(code []byte) Action {
	return Action{
		Enum: borsh.Enum(OrdDeployContract),
		DeployContract: ActionDeployContract{
			Code: code,
		},
	}
}

type ActionFunctionCall struct {
	MethodName string  `json:"method_name"`
	Args       []byte  `json:"args"`
	Gas        Gas     `json:"gas"`
	Deposit    Balance `json:"deposit"`
}

func NewFunctionCall(methodName string, args []byte, gas Gas, deposit Balance) Action {
	return Action{
		Enum: borsh.Enum(OrdFunctionCall),
		FunctionCall: ActionFunctionCall{
			MethodName: methodName,
			Args:       args,
			Gas:        gas,
			Deposit:    deposit,
		},
	}
}

func (f ActionFunctionCall) String() string {
	return fmt.Sprintf("FunctionCall{MethodName: %s, Args: %s, Gas: %d, Deposit: %s}", f.MethodName, f.Args, f.Gas, f.Deposit)
}

type ActionTransfer struct {
	Deposit Balance `json:"deposit"`
}

func NewTransfer(deposit Balance) Action {
	return Action{
		Enum: borsh.Enum(OrdTransfer),
		Transfer: ActionTransfer{
			Deposit: deposit,
		},
	}
}

func (t ActionTransfer) String() string {
	return fmt.Sprintf("Transfer{Deposit: %s}", t.Deposit)
}

type ActionStake struct {
	// Amount of tokens to stake.
	Stake Balance `json:"stake"`
	// Validator key which will be used to sign transactions on behalf of singer_id
	PublicKey PublicKey `json:"public_key"`
}

func NewStake(stake Balance, publicKey PublicKey) Action {
	return Action{
		Enum: borsh.Enum(OrdStake),
		Stake: ActionStake{
			Stake:     stake,
			PublicKey: publicKey,
		},
	}
}

type ActionAddKey struct {
	PublicKey PublicKey             `json:"public_key"`
	AccessKey ActionAddKeyAccessKey `json:"access_key"`
}

func NewAddKey(publicKey PublicKey, nonce Nonce, permission AccessKeyPermission) Action {
	return Action{
		Enum:   borsh.Enum(OrdAddKey),
		AddKey: ActionAddKey{},
	}
}

type ActionAddKeyAccessKey struct {
	Nonce      Nonce               `json:"nonce"`
	Permission AccessKeyPermission `json:"permission"`
}

type JsonActionAddKey struct {
	PublicKey Base58PublicKey       `json:"public_key"`
	AccessKey ActionAddKeyAccessKey `json:"access_key"`
}

func (a ActionAddKey) MarshalJSON() (b []byte, err error) {
	v := JsonActionAddKey{
		PublicKey: a.PublicKey.ToBase58PublicKey(),
		AccessKey: a.AccessKey,
	}
	b, err = json.Marshal(&v)
	return
}

func (a *ActionAddKey) UnmarshalJSON(b []byte) (err error) {
	var v JsonActionAddKey
	if err = json.Unmarshal(b, &v); err != nil {
		return
	}

	*a = ActionAddKey{
		PublicKey: v.PublicKey.ToPublicKey(),
		AccessKey: v.AccessKey,
	}

	return
}

type ActionDeleteKey struct {
	PublicKey PublicKey `json:"public_key"`
}

func NewDeleteKey(publicKey PublicKey) Action {
	return Action{
		Enum: borsh.Enum(OrdDeleteKey),
		DeleteKey: ActionDeleteKey{
			PublicKey: publicKey,
		},
	}
}

type JsonActionDeleteKey struct {
	PublicKey Base58PublicKey `json:"public_key"`
}

func (a ActionDeleteKey) MarshalJSON() (b []byte, err error) {
	v := JsonActionDeleteKey{
		PublicKey: a.PublicKey.ToBase58PublicKey(),
	}
	b, err = json.Marshal(&v)
	return
}

func (a *ActionDeleteKey) UnmarshalJSON(b []byte) (err error) {
	var v JsonActionDeleteKey
	if err = json.Unmarshal(b, &v); err != nil {
		return
	}

	*a = ActionDeleteKey{
		PublicKey: v.PublicKey.ToPublicKey(),
	}

	return
}

type ActionDeleteAccount struct {
	BeneficiaryID AccountID `json:"beneficiary_id"`
}

func NewDeleteAccount(beneficiaryID AccountID) Action {
	return Action{
		Enum: borsh.Enum(OrdDeleteAccount),
		DeleteAccount: ActionDeleteAccount{
			BeneficiaryID: beneficiaryID,
		},
	}
}
