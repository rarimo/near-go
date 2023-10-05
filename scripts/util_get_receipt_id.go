package scripts

import (
	"encoding/json"
	"github.com/rarimo/near-go/common"
	"strings"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	ReceiptNotFoundError = errors.New("receipt not found")
	eventPrefix          = "EVENT_JSON:"
)

func GetDepositedReceiptID(tx common.FinalExecutionOutcomeView, eventType common.LogEventType, bridge string, token *string, tokenID *string, amount *common.Balance) (*common.Hash, error) {
	var receiptID *common.Hash

	for _, receipt := range tx.ReceiptsOutcome {
		if receipt.Outcome.ExecutorID != bridge {
			continue
		}
		if len(receipt.Outcome.Logs) == 0 {
			return nil, errors.Wrap(ReceiptNotFoundError, "receipt has no logs")
		}

		for _, outcome := range receipt.Outcome.Logs {
			rawLog := strings.ReplaceAll(outcome, eventPrefix, "")

			var log common.LogEvent
			err := json.Unmarshal([]byte(rawLog), &log)
			if err != nil {
				return nil, errors.Wrap(err, "failed to unmarshal event")
			}

			if log.Event != eventType {
				continue
			}
			if len(log.Data) == 0 {
				return nil, errors.Wrap(ReceiptNotFoundError, "receipt has no data")
			}

			// bridge log always will have only one data object
			logData := log.Data[0]

			if token != nil && logData.Token != nil && *token != *logData.Token {
				continue
			}
			if tokenID != nil && logData.TokenID != nil && *tokenID != *logData.TokenID {
				continue
			}
			if amount != nil && logData.Amount != nil && !amount.Equals(*logData.Amount) {
				continue
			}

			return &receipt.ID, nil
		}
	}

	if receiptID == nil {
		return nil, ReceiptNotFoundError
	}

	return receiptID, nil
}
