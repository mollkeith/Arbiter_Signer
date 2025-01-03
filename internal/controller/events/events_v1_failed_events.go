package events

import (
	"context"
	"os"
	"strings"

	v1 "github.com/BeL2Labs/Arbiter_Signer/api/events/v1"
	"github.com/BeL2Labs/Arbiter_Signer/internal/consts"
	"github.com/BeL2Labs/Arbiter_Signer/utility/contract_abi"
	"github.com/BeL2Labs/Arbiter_Signer/utility/events"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/frame/g"
)

func (c *ControllerV1) FailedEvents(ctx context.Context, req *v1.FailedEventsReq) (res *v1.FailedEventsRes, err error) {
	loanABI, err := abi.JSON(strings.NewReader(contract_abi.ArbiterABI))
	if err != nil {
		return nil, err
	}

	filedFilePath := getExpandedPath(consts.FailedEventFilePath)
	failedFiles, err := os.ReadDir(filedFilePath)
	if err != nil {
		return nil, err
	}

	evs := make([]events.EventInfo, 0)
	for _, file := range failedFiles {
		filePath := filedFilePath + "/" + file.Name()
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			g.Log().Error(ctx, "read file:", filePath, "error:", err)
			continue
		}
		logEvt, err := decodeLogEvtByFileContent(fileContent)
		if err != nil {
			g.Log().Error(ctx, "decodeLogEvtByFileContent failed, file:", filePath, "err:", err)
			continue
		}
		var ev = make(map[string]interface{})
		err = loanABI.UnpackIntoMap(ev, "ArbitrationRequested", logEvt.EventData)
		if err != nil {
			g.Log().Error(ctx, "UnpackIntoMap file:", filePath, "error:", err)
			continue
		}
		g.Log().Info(ctx, "ev", ev)
		dappAddress := logEvt.Topics[1]
		queryId := logEvt.Topics[2]
		// rawData := ev["btcTx"].([]byte)
		// script := ev["script"].([]byte)
		arbitratorAddress := ev["arbitrator"].(common.Address)

		evs = append(evs, events.EventInfo{
			EventName:         "ArbitrationRequested",
			EventID:           logEvt.TxHash.String(),
			QueryID:           queryId.String(),
			Block:             logEvt.Block,
			ArbitratorAddress: arbitratorAddress.String(),
			DappAddress:       dappAddress.String(),
			Status:            "failed",
		})

	}

	res = &v1.FailedEventsRes{
		Events: evs,
	}
	return res, nil
}
