package events

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/BeL2Labs/Arbiter_Signer/api/events/v1"
	"github.com/BeL2Labs/Arbiter_Signer/internal/consts"
	"github.com/BeL2Labs/Arbiter_Signer/utility/contract_abi"
	"github.com/BeL2Labs/Arbiter_Signer/utility/events"
)

func (c *ControllerV1) AllEvents(ctx context.Context, req *v1.AllEventsReq) (res *v1.AllEventsRes, err error) {
	loanABI, err := abi.JSON(strings.NewReader(contract_abi.ArbiterABI))
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}

	evs := make([]events.EventInfo, 0)
	filedFilePath := getExpandedPath(consts.FailedEventFilePath)
	failedFiles, err := os.ReadDir(filedFilePath)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}
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

	succeedFilePath := getExpandedPath(consts.SucceedEventFilePath)
	succeedFiles, err := os.ReadDir(succeedFilePath)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}
	for _, file := range succeedFiles {
		filePath := succeedFilePath + "/" + file.Name()
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

		dappAddress := logEvt.Topics[1]
		queryId := logEvt.Topics[2]
		arbitratorAddress := ev["arbitrator"].(common.Address)
		evs = append(evs, events.EventInfo{
			EventName:         "ArbitrationRequested",
			EventID:           logEvt.TxHash.String(),
			QueryID:           queryId.String(),
			Block:             logEvt.Block,
			ArbitratorAddress: arbitratorAddress.String(),
			DappAddress:       dappAddress.String(),
			Status:            "success",
		})
	}

	requiredFilePath := getExpandedPath(consts.RequestEventFilePath)
	requiredFiles, err := os.ReadDir(requiredFilePath)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}
	for _, file := range requiredFiles {
		filePath := requiredFilePath + "/" + file.Name()
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
		dappAddress := logEvt.Topics[1]
		queryId := logEvt.Topics[2]
		arbitratorAddress := ev["arbitrator"].(common.Address)
		evs = append(evs, events.EventInfo{
			EventName:         "ArbitrationRequested",
			EventID:           logEvt.TxHash.String(),
			QueryID:           queryId.String(),
			Block:             logEvt.Block,
			ArbitratorAddress: arbitratorAddress.String(),
			DappAddress:       dappAddress.String(),
			Status:            "required",
		})
	}

	res = &v1.AllEventsRes{
		Events: evs,
	}

	return res, nil
}

func decodeLogEvtByFileContent(content []byte) (*events.ContractLogEvent, error) {
	logEvt := &events.ContractLogEvent{}
	err := gob.NewDecoder(bytes.NewReader(content)).Decode(logEvt)
	if err != nil {
		return nil, err
	}
	return logEvt, nil
}

func getExpandedPath(path string) string {
	if len(path) > 0 && path[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return path
		}
		path = filepath.Join(homeDir, path[2:])
	}
	return path
}
