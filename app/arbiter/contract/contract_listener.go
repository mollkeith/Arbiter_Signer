// Copyright (c) 2025 The bel2 developers

package contract

import (
	"context"
	"math"
	"math/big"

	"Arbiter_Signer/app/arbiter/contract/events"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
)

type ContractListener struct {
	loanContract  common.Address
	queryClient   *CrossClient
	listeneTopics []common.Hash
	ctx           context.Context
	chan_events   chan *events.ContractLogEvent
}

func NewListener(ctx context.Context, client *CrossClient,
	loanContract common.Address, chan_event chan *events.ContractLogEvent) (*ContractListener, error) {
	c := &ContractListener{
		queryClient:  client,
		loanContract: loanContract,
		ctx:          ctx,
		chan_events:  chan_event,
	}
	c.listeneTopics = make([]common.Hash, 0)
	return c, nil
}

func (c *ContractListener) Start(startHeight uint64) (uint64, error) {
	endBlock, err := c.queryClient.GetLatestHeight()
	if err != nil {
		g.Log().Warning(c.ctx, "GetLatestHeight failed", err)
		return math.MaxUint64, err
	}
	endBlock -= 2

	distance := uint64(10000)
	var toBlock = startHeight
	loanQuery := c.queryClient.BuildQuery(c.loanContract, c.listeneTopics, nil, nil)
	for i := startHeight; i <= endBlock; i = toBlock + 1 {
		if i+distance < endBlock {
			toBlock = i + distance
		} else {
			toBlock = endBlock
		}
		loanQuery.FromBlock = big.NewInt(0).SetUint64(i)
		loanQuery.ToBlock = big.NewInt(0).SetUint64(toBlock)
		g.Log().Infof(c.ctx, "pull block from %v to %v", i, toBlock)
		err = c.filterLoanEvent(loanQuery)
		if err != nil {
			g.Log().Error(c.ctx, "filter filterLoanEvent failed, error:", err)
		}
	}

	return endBlock, nil
}

func (c *ContractListener) filterLoanEvent(query ethereum.FilterQuery) error {
	logs, err := c.queryClient.FilterLogs(c.ctx, query)
	if err != nil {
		return err
	}
	for _, l := range logs {
		evt := &events.ContractLogEvent{
			EventData: l.Data,
			TxHash:    l.TxHash,
			Topics:    l.Topics,
			Block:     l.BlockNumber,
			TxIndex:   l.TxIndex,
		}
		c.chan_events <- evt
	}
	return nil
}