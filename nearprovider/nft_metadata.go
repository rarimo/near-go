package nearprovider

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/constants"
	"github.com/rarimo/near-go/nearclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *provider) GetNFTMetadata(ctx context.Context, token common.AccountID, tokenID string) (*common.NftMetadataView, error) {
	resp, err := p.c.ContractViewCallFunction(
		ctx,
		token,
		constants.ContractNftGet,
		base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("{\"token_id\":\"%s\"}", tokenID))),
		nearclient.FinalityFinal(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token metadata")
	}

	var result *common.NftView
	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal nft view")
	}

	return result.Metadata, nil
}
