package tests

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/scripts"
	"testing"
)

var response = json.RawMessage(`{
        "receipts_outcome": [
            {
                "block_hash": "7Qs3KgV5KnFE6Cb9zvF4Qmuy35jgtQtSQbVQiLCS3gy1",
                "id": "AhQ2yUg3zbpwhh42ZLEXdnJCmNQfvGeC5cLVr6YKpkWo",
                "outcome": {
                    "executor_id": "non_fungible_token_original.napalmpapalam.testnet",
                    "gas_burnt": 19435210535617,
                    "logs": [
                        "EVENT_JSON:{\"standard\":\"nep171\",\"version\":\"1.0.0\",\"event\":\"nft_transfer\",\"data\":[{\"old_owner_id\":\"napalmpapalam.testnet\",\"new_owner_id\":\"bridge.napalmpapalam.testnet\",\"token_ids\":[\"1\"]}]}",
                        "EVENT_JSON:{\"standard\":\"nep171\",\"version\":\"1.0.0\",\"event\":\"nft_transfer\",\"data\":[{\"old_owner_id\":\"napalmpapalam.testnet\",\"new_owner_id\":\"bridge.napalmpapalam.testnet\",\"token_ids\":[\"2\"]}]}"
                    ],
                    "metadata": {},
                    "receipt_ids": [
                        "HvR7Z6U4qdJP4eE5A9LrFDPKmr5f4pA9pnJF1CsfhfT9",
                        "4teub4Vg3pT7FUziTFWHwAd5u1SLUeVCZhRvNQrq8RHq",
                        "75ZRuevHoM9LBp7XxrcnJtJ2fuiLj9sULyDuuVg5rZtQ",
                        "5v9gdzzRj7rxo3v7jSFQheH52HYFwywjwHSbgqharMfW",
                        "Go4MXz6pHrA2HLF2Wm5SeKiT9RBgqkBon9NMspKJhy3T"
                    ],
                    "status": {
                        "SuccessReceiptId": "5v9gdzzRj7rxo3v7jSFQheH52HYFwywjwHSbgqharMfW"
                    },
                    "tokens_burnt": "1943521053561700000000"
                },
                "proof": []
            },
            {
                "block_hash": "7AabLjyNZeHevosJWFB1jtRNKyzWV73vADP34TwFEPme",
                "id": "HvR7Z6U4qdJP4eE5A9LrFDPKmr5f4pA9pnJF1CsfhfT9",
                "outcome": {
                    "executor_id": "bridge.napalmpapalam.testnet",
                    "gas_burnt": 3190610576685,
                    "logs": [
                        "EVENT_JSON:{\"standard\":\"nep171\",\"version\":\"1.0.0\",\"event\":\"nft_deposited\",\"data\":[{\"token\":\"non_fungible_token_original.napalmpapalam.testnet\",\"token_id\":\"1\",\"receiver\":\"napalmpapalam.testnet\",\"chain\":\"Near\",\"is_wrapped\":false}]}"
                    ],
                    "metadata": {},
                    "receipt_ids": [
                        "9ZLxGFUVTEQSYfWgxZJHkHMDpk5Rfa4pH9ZdnYvLo83x"
                    ],
                    "status": {
                        "SuccessValue": "ZmFsc2U="
                    },
                    "tokens_burnt": "319061057668500000000"
                },
                "proof": []
            },
            {
                "block_hash": "7cJe4WbeLpqDnR5D5KFoBpCPEe8xXMM9w1Mvru2a7jyG",
                "id": "9ZLxGFUVTEQSYfWgxZJHkHMDpk5Rfa4pH9ZdnYvLo83x",
                "outcome": {
                    "executor_id": "napalmpapalam.testnet",
                    "gas_burnt": 223182562500,
                    "logs": [],
                    "metadata": {
                        "gas_profile": [],
                        "version": 1
                    },
                    "receipt_ids": [],
                    "status": {
                        "SuccessValue": ""
                    },
                    "tokens_burnt": "0"
                },
                "proof": []
            },
            {
                "block_hash": "7cJe4WbeLpqDnR5D5KFoBpCPEe8xXMM9w1Mvru2a7jyG",
                "id": "4teub4Vg3pT7FUziTFWHwAd5u1SLUeVCZhRvNQrq8RHq",
                "outcome": {
                    "executor_id": "non_fungible_token_original.napalmpapalam.testnet",
                    "gas_burnt": 2947052313301,
                    "logs": [],
                    "metadata": {},
                    "receipt_ids": [
                        "5QqH1Qp88YpGNzHkdQdEoLchLssvd4Yg8N6MtX9L3SWv"
                    ],
                    "status": {
                        "SuccessValue": "dHJ1ZQ=="
                    },
                    "tokens_burnt": "294705231330100000000"
                },
                "proof": []
            },
            {
                "block_hash": "ABc8895GDrDQyoZsSDQDPd27EfMEuzsiZDr4z3zLfffE",
                "id": "5QqH1Qp88YpGNzHkdQdEoLchLssvd4Yg8N6MtX9L3SWv",
                "outcome": {
                    "executor_id": "napalmpapalam.testnet",
                    "gas_burnt": 223182562500,
                    "logs": [],
                    "metadata": {
                        "gas_profile": [],
                        "version": 1
                    },
                    "receipt_ids": [],
                    "status": {
                        "SuccessValue": ""
                    },
                    "tokens_burnt": "0"
                },
                "proof": []
            },
            {
                "block_hash": "7AabLjyNZeHevosJWFB1jtRNKyzWV73vADP34TwFEPme",
                "id": "75ZRuevHoM9LBp7XxrcnJtJ2fuiLj9sULyDuuVg5rZtQ",
                "outcome": {
                    "executor_id": "bridge.napalmpapalam.testnet",
                    "gas_burnt": 2802309502239,
                    "logs": [
                        "EVENT_JSON:{\"standard\":\"nep171\",\"version\":\"1.0.0\",\"event\":\"nft_deposited\",\"data\":[{\"token\":\"non_fungible_token_original.napalmpapalam.testnet\",\"token_id\":\"2\",\"receiver\":\"napalmpapalam.testnet\",\"chain\":\"Near\",\"is_wrapped\":false}]}"
                    ],
                    "metadata": {},
                    "receipt_ids": [
                        "DuqmiMF8EC3C7shpjZVyvf8rbf6J9a9dd3FMDr5MciZm"
                    ],
                    "status": {
                        "SuccessValue": "ZmFsc2U="
                    },
                    "tokens_burnt": "280230950223900000000"
                },
                "proof": []
            },
            {
                "block_hash": "7cJe4WbeLpqDnR5D5KFoBpCPEe8xXMM9w1Mvru2a7jyG",
                "id": "DuqmiMF8EC3C7shpjZVyvf8rbf6J9a9dd3FMDr5MciZm",
                "outcome": {
                    "executor_id": "napalmpapalam.testnet",
                    "gas_burnt": 223182562500,
                    "logs": [],
                    "metadata": {
                        "gas_profile": [],
                        "version": 1
                    },
                    "receipt_ids": [],
                    "status": {
                        "SuccessValue": ""
                    },
                    "tokens_burnt": "0"
                },
                "proof": []
            },
            {
                "block_hash": "7cJe4WbeLpqDnR5D5KFoBpCPEe8xXMM9w1Mvru2a7jyG",
                "id": "5v9gdzzRj7rxo3v7jSFQheH52HYFwywjwHSbgqharMfW",
                "outcome": {
                    "executor_id": "non_fungible_token_original.napalmpapalam.testnet",
                    "gas_burnt": 2705948842189,
                    "logs": [],
                    "metadata": {},
                    "receipt_ids": [
                        "DdMgPycLY4Chkz7avwg65kccGQ5K8Anx1EZEPSoPwKJP"
                    ],
                    "status": {
                        "SuccessValue": "dHJ1ZQ=="
                    },
                    "tokens_burnt": "270594884218900000000"
                },
                "proof": []
            },
            {
                "block_hash": "ABc8895GDrDQyoZsSDQDPd27EfMEuzsiZDr4z3zLfffE",
                "id": "DdMgPycLY4Chkz7avwg65kccGQ5K8Anx1EZEPSoPwKJP",
                "outcome": {
                    "executor_id": "napalmpapalam.testnet",
                    "gas_burnt": 223182562500,
                    "logs": [],
                    "metadata": {
                        "gas_profile": [],
                        "version": 1
                    },
                    "receipt_ids": [],
                    "status": {
                        "SuccessValue": ""
                    },
                    "tokens_burnt": "0"
                },
                "proof": []
            },
            {
                "block_hash": "7AabLjyNZeHevosJWFB1jtRNKyzWV73vADP34TwFEPme",
                "id": "Go4MXz6pHrA2HLF2Wm5SeKiT9RBgqkBon9NMspKJhy3T",
                "outcome": {
                    "executor_id": "napalmpapalam.testnet",
                    "gas_burnt": 223182562500,
                    "logs": [],
                    "metadata": {
                        "gas_profile": [],
                        "version": 1
                    },
                    "receipt_ids": [],
                    "status": {
                        "SuccessValue": ""
                    },
                    "tokens_burnt": "0"
                },
                "proof": []
            }
        ],
        "status": {
            "SuccessValue": "dHJ1ZQ=="
        },
        "transaction": {
            "actions": [
                {
                    "FunctionCall": {
                        "args": "eyJyZWNlaXZlcl9pZCI6ImJyaWRnZS5uYXBhbG1wYXBhbGFtLnRlc3RuZXQiLCJ0b2tlbl9pZCI6IjEiLCJtc2ciOiJ7XCJ0b2tlblwiOlwibm9uX2Z1bmdpYmxlX3Rva2VuX29yaWdpbmFsLm5hcGFsbXBhcGFsYW0udGVzdG5ldFwiLFwicmVjZWl2ZXJcIjpcIm5hcGFsbXBhcGFsYW0udGVzdG5ldFwiLFwiY2hhaW5cIjpcIk5lYXJcIixcImlzX3dyYXBwZWRcIjpmYWxzZX0ifQ==",
                        "deposit": "1",
                        "gas": 150000000000000,
                        "method_name": "nft_transfer_call"
                    }
                },
                {
                    "FunctionCall": {
                        "args": "eyJyZWNlaXZlcl9pZCI6ImJyaWRnZS5uYXBhbG1wYXBhbGFtLnRlc3RuZXQiLCJ0b2tlbl9pZCI6IjIiLCJtc2ciOiJ7XCJ0b2tlblwiOlwibm9uX2Z1bmdpYmxlX3Rva2VuX29yaWdpbmFsLm5hcGFsbXBhcGFsYW0udGVzdG5ldFwiLFwicmVjZWl2ZXJcIjpcIm5hcGFsbXBhcGFsYW0udGVzdG5ldFwiLFwiY2hhaW5cIjpcIk5lYXJcIixcImlzX3dyYXBwZWRcIjpmYWxzZX0ifQ==",
                        "deposit": "1",
                        "gas": 150000000000000,
                        "method_name": "nft_transfer_call"
                    }
                }
            ],
            "hash": "8QodxzYVWgBGBC93KYmerEepGAufNs38MvBJjDrbHQ8u",
            "nonce": 101318032000295,
            "public_key": "ed25519:BhTuMWScWqoyuzP7JkXGq2eZqZ7KX8jToF6W2Vjmummn",
            "receiver_id": "non_fungible_token_original.napalmpapalam.testnet",
            "signature": "ed25519:3VoMgHznvfPrA5YmwMU1T2n9UrfxDFRzLd1c17QDSVtcFTWySe8d3RR8jtdJ1TcVS1xE5wxotUoubVqJXcze8eEX",
            "signer_id": "napalmpapalam.testnet"
        },
        "transaction_outcome": {
            "block_hash": "8P6kb38xZ8YMQ4LmNPrDK2ai9omPgu74MVRfcTC6n82R",
            "id": "8QodxzYVWgBGBC93KYmerEepGAufNs38MvBJjDrbHQ8u",
            "outcome": {
                "executor_id": "napalmpapalam.testnet",
                "gas_burnt": 4748815501508,
                "logs": [],
                "metadata": {
                    "gas_profile": null,
                    "version": 1
                },
                "receipt_ids": [
                    "AhQ2yUg3zbpwhh42ZLEXdnJCmNQfvGeC5cLVr6YKpkWo"
                ],
                "status": {
                    "SuccessReceiptId": "AhQ2yUg3zbpwhh42ZLEXdnJCmNQfvGeC5cLVr6YKpkWo"
                },
                "tokens_burnt": "474881550150800000000"
            },
            "proof": []
        }
    }`)

func TestUtilGetReceiptId(t *testing.T) {
	var resp client.FinalExecutionOutcomeView
	err := json.Unmarshal(response, &resp)
	if !assert.NoError(t, err) {
		return
	}

	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	receiptId, err := scripts.GetDepositedReceiptID(
		resp,
		client.LogEventTypeNftDeposited,
		cfg.BridgeAddress,
		&cfg.NftAddressOriginal,
		&cfg.TokenID,
		nil,
	)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "HvR7Z6U4qdJP4eE5A9LrFDPKmr5f4pA9pnJF1CsfhfT9", receiptId.String())
}
