// Copyright (c) 2025 The bel2 developers

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EventLogReq struct {
	g.Meta `path:"/event_log" tags:"Event log" method:"get" summary:"Get all event logs"`
}

type EventLogRes struct {
	EventLog string `json:"eventLog"`
}
