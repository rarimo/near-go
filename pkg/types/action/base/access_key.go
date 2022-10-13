package base

import (
	"encoding/json"
	"fmt"
	"github.com/eteu-technologies/borsh-go"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
)

type AccessKeyPermission struct {
	Enum borsh.Enum `borsh_enum:"true"`

	FunctionCallPermission AccessKeyFunctionCallPermission
	FullAccessPermission   struct{}
}

type fullAccessPermissionWrapper struct {
	FunctionCall AccessKeyFunctionCallPermission `json:"FunctionCall"`
}

func NewFunctionCallPermission(allowance types.Balance, receiverID types.AccountID, methodNames []string) AccessKeyPermission {
	return AccessKeyPermission{
		Enum: borsh.Enum(0),
		FunctionCallPermission: AccessKeyFunctionCallPermission{
			Allowance:   &allowance,
			ReceiverID:  receiverID,
			MethodNames: methodNames,
		},
	}
}

func NewFunctionCallUnlimitedAllowancePermission(receiverID types.AccountID, methodNames []string) AccessKeyPermission {
	return AccessKeyPermission{
		Enum: borsh.Enum(0),
		FunctionCallPermission: AccessKeyFunctionCallPermission{
			Allowance:   nil,
			ReceiverID:  receiverID,
			MethodNames: methodNames,
		},
	}
}

func NewFullAccessPermission() AccessKeyPermission {
	return AccessKeyPermission{
		Enum: borsh.Enum(1),
	}
}

func (a AccessKeyPermission) MarshalJSON() ([]byte, error) {
	if a.IsFullAccess() {
		return []byte(`"FullAccess"`), nil
	}

	var v fullAccessPermissionWrapper
	v.FunctionCall = a.FunctionCallPermission

	return json.Marshal(&v)
}

func (a *AccessKeyPermission) UnmarshalJSON(b []byte) error {
	if len(b) > 0 && b[0] == '{' {
		var permission fullAccessPermissionWrapper
		if err := json.Unmarshal(b, &permission); err != nil {
			return err
		}

		*a = AccessKeyPermission{
			Enum:                   borsh.Enum(0),
			FunctionCallPermission: permission.FunctionCall,
		}
		return nil
	}

	var value string

	if value == "FullAccess" {
		*a = NewFullAccessPermission()
		return nil
	}

	return errors.New(fmt.Sprintf("unknown permission '%s'", value))
}

func (a AccessKeyPermission) String() string {
	var value = "FullAccess"
	if a.IsFunctionCall() {
		value = a.FunctionCallPermission.String()
	}
	return fmt.Sprintf("AccessKeyPermission{%s}", value)
}

func (a *AccessKeyPermission) IsFunctionCall() bool {
	return a.Enum == 0
}

func (a *AccessKeyPermission) IsFullAccess() bool {
	return a.Enum == 1
}

type AccessKeyFunctionCallPermission struct {
	Allowance   *types.Balance  `json:"allowance"`
	ReceiverID  types.AccountID `json:"receiver_id"`
	MethodNames []string        `json:"method_names"`
}

func (a AccessKeyFunctionCallPermission) String() string {
	return fmt.Sprintf("AccessKeyFunctionCallPermission{Allowance=%v, ReceiverID=%v, MethodNames=%v}", a.Allowance, a.ReceiverID, a.MethodNames)
}
