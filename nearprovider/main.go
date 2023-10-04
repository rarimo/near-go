package nearprovider

import (
	"context"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/nearclient"
	"github.com/rarimo/near-go/nearclient/models"
	"github.com/rarimo/near-go/nearprovider/s3"
	"gitlab.com/distributed_lab/logan/v3"
)

type Provider interface {
	ListBlocks(ctx context.Context, limit uint64, fromBlock common.BlockHeight) ([]common.BlockHeight, error)
	GetMessage(ctx context.Context, height common.BlockHeight) (*Message, error)
	GetTransaction(ctx context.Context, hash common.Hash, sender common.AccountID) (*models.FinalExecutionOutcomeWithReceiptView, error)
	GetLastKnownBlockHeight(ctx context.Context) (*common.BlockHeight, error)
	GetNFTMetadata(ctx context.Context, token common.AccountID, tokenID string) (*common.NftMetadataView, error)
}

type provider struct {
	c   *nearclient.Client
	hc  *nearclient.Client
	s3  s3.Connector
	log *logan.Entry
}

func New(client, historyClient *nearclient.Client, s3 s3.Connector, log *logan.Entry) Provider {
	return &provider{
		client,
		historyClient,
		s3,
		log,
	}
}
