package nearprovider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/nearclient/models"
	"github.com/rarimo/near-go/nearprovider/s3"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strconv"
	"strings"
)

type Message struct {
	Block  *models.BlockView   `json:"block"`
	Shards []*models.ShardView `json:"shards"`
}

var ErrNoShardsAvailable = errors.New("No shards available")

// ListBlocks - queries the list of the objects in the bucket, grouped by "/" delimiter.
// Returns the list of blocks that can be fetched
// See more about data structure https://github.com/near/near-lake#data-structure
func (p *provider) ListBlocks(ctx context.Context, limit uint64, fromBlock common.BlockHeight) ([]common.BlockHeight, error) {
	res, err := p.s3.ListObjects(ctx, int64(limit), normalizeBlockHeight(fromBlock))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get list blocks")
	}

	blocks := make([]common.BlockHeight, 0)

	for _, block := range res.CommonPrefixes {
		if block.Prefix == nil {
			continue
		}

		parsed, err := strconv.ParseUint(strings.TrimSuffix(*block.Prefix, s3.Delimiter), 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse block height into unit64")
		}

		blocks = append(blocks, parsed)
	}

	p.log.WithFields(logan.F{
		"start_from_block": fromBlock,
		"batch_size":       len(blocks),
	}).Debug("Listed blocks")

	return blocks, nil
}

func (p *provider) getBlock(ctx context.Context, blockHeight common.BlockHeight) (*models.BlockView, error) {
	res, err := p.s3.GetObject(ctx, fmt.Sprintf("%s%sblock.json", normalizeBlockHeight(blockHeight), s3.Delimiter))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block")
	}

	var block models.BlockView

	err = json.NewDecoder(res.Body).Decode(&block)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse block response")
	}

	p.log.WithFields(logan.F{
		"block_height": blockHeight,
	}).Debug("Got block")

	return &block, nil
}

func (p *provider) getShard(ctx context.Context, blockHeight common.BlockHeight, shardID uint64) (*models.ShardView, error) {
	res, err := p.s3.GetObject(ctx, fmt.Sprintf("%s%sshard_%d.json", normalizeBlockHeight(blockHeight), s3.Delimiter, shardID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shard")
	}

	var shard models.ShardView
	err = json.NewDecoder(res.Body).Decode(&shard)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse shard response")
	}

	p.log.WithFields(logan.F{
		"shard_id":     shardID,
		"block_height": blockHeight,
	}).Debug("Got shard")

	return &shard, nil
}

func (p *provider) getShards(ctx context.Context, blockHeight common.BlockHeight, numberOfShards uint64) ([]*models.ShardView, error) {
	if numberOfShards == 0 {
		return nil, ErrNoShardsAvailable
	}

	shards := make([]*models.ShardView, numberOfShards)

	for i := uint64(0); i < numberOfShards; i++ {
		shard, err := p.getShard(ctx, blockHeight, i)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get shard")
		}
		shards[i] = shard
	}

	p.log.WithFields(logan.F{
		"block_height": blockHeight,
		"shards_count": len(shards),
	}).Debug("Got shards")

	return shards, nil
}

// GetMessage by the given block height gets the objects:
// - block.json
// - shard_N.json
// Returns the result as `Message` struct
func (p *provider) GetMessage(ctx context.Context, blockHeight common.BlockHeight) (*Message, error) {
	block, err := p.getBlock(ctx, blockHeight)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block")
	}

	shards, err := p.getShards(ctx, blockHeight, uint64(len(block.Chunks)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shards")
	}

	return &Message{
		Block:  block,
		Shards: shards,
	}, nil
}

func normalizeBlockHeight(blockHeight common.BlockHeight) string {
	return fmt.Sprintf("%012d", blockHeight)
}
