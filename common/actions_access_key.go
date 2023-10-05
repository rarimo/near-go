package common

import (
	"encoding/json"
	"fmt"
	"github.com/near/borsh-go"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strings"
)

type ActionAccessKeyPermission struct {
	Enum borsh.Enum `borsh_enum:"true"`

	FunctionCallPermission AccessKeyFunctionCallPermission
	FullAccessPermission   struct{}
}

func NewFunctionCallPermission(allowance Balance, receiverID AccountID, methodNames []string) ActionAccessKeyPermission {
	return ActionAccessKeyPermission{
		Enum: borsh.Enum(0),
		FunctionCallPermission: AccessKeyFunctionCallPermission{
			Allowance:   &allowance,
			ReceiverID:  receiverID,
			MethodNames: methodNames,
		},
	}
}

func NewFunctionCallUnlimitedAllowancePermission(receiverID AccountID, methodNames []string) ActionAccessKeyPermission {
	return ActionAccessKeyPermission{
		Enum: borsh.Enum(0),
		FunctionCallPermission: AccessKeyFunctionCallPermission{
			Allowance:   nil,
			ReceiverID:  receiverID,
			MethodNames: methodNames,
		},
	}
}

func NewFullAccessPermission() ActionAccessKeyPermission {
	return ActionAccessKeyPermission{
		Enum: borsh.Enum(1),
	}
}

type FullAccessPermissionWrapper struct {
	FunctionCall AccessKeyFunctionCallPermission `json:"FunctionCall"`
}

func (a ActionAccessKeyPermission) MarshalJSON() ([]byte, error) {
	if a.IsFullAccess() {
		return []byte(`"FullAccess"`), nil
	}

	var v FullAccessPermissionWrapper
	v.FunctionCall = a.FunctionCallPermission

	return json.Marshal(&v)
}

func (a *ActionAccessKeyPermission) UnmarshalJSON(b []byte) error {
	if len(b) > 0 && b[0] == '{' {
		var permission FullAccessPermissionWrapper
		if err := json.Unmarshal(b, &permission); err != nil {
			return err
		}

		*a = ActionAccessKeyPermission{
			Enum:                   borsh.Enum(0),
			FunctionCallPermission: permission.FunctionCall,
		}
		return nil
	}

	value := strings.ReplaceAll(string(b), "\"", "")

	if value == "FullAccess" {
		*a = NewFullAccessPermission()
		return nil
	}

	return errors.New(fmt.Sprintf("unknown permission '%s'", value))
}

func (a ActionAccessKeyPermission) String() string {
	var value = "FullAccess"
	if a.IsFunctionCall() {
		value = a.FunctionCallPermission.String()
	}
	return fmt.Sprintf("AccessKeyPermission{%s}", value)
}

func (a *ActionAccessKeyPermission) IsFunctionCall() bool {
	return a.Enum == 0
}

func (a *ActionAccessKeyPermission) IsFullAccess() bool {
	return a.Enum == 1
}

type AccessKeyFunctionCallPermission struct {
	Allowance   *Balance  `json:"allowance"`
	ReceiverID  AccountID `json:"receiver_id"`
	MethodNames []string  `json:"method_names"`
}

func (a AccessKeyFunctionCallPermission) String() string {
	return fmt.Sprintf("AccessKeyFunctionCallPermission{Allowance=%v, ReceiverID=%v, MethodNames=%v}", a.Allowance, a.ReceiverID, a.MethodNames)
}
