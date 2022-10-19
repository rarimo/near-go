package scripts

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/hash"
	"strings"
)

var (
	ReceiptNotFoundError = errors.New("receipt not found")
	eventPrefix          = "EVENT_JSON:"
)

func GetDepositedReceiptID(tx client.FinalExecutionOutcomeView, eventType client.LogEventType, bridge string, token string, tokenID *string, amount *types.Balance) (*hash.CryptoHash, error) {
	var receiptID *hash.CryptoHash

	for _, receipt := range tx.ReceiptsOutcome {
		if receipt.Outcome.ExecutorID != bridge {
			continue
		}
		if len(receipt.Outcome.Logs) == 0 {
			return nil, errors.Wrap(ReceiptNotFoundError, "receipt has no logs")
		}

		for _, outcome := range receipt.Outcome.Logs {
			rawLog := strings.ReplaceAll(outcome, eventPrefix, "")

			var log client.LogEvent
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

			if logData.Token != token {
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
