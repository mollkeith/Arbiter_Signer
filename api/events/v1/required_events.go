// Copyright (c) 2025 The bel2 developers

package v1

import (
	"github.com/BeL2Labs/Arbiter_Signer/utility/events"

	"github.com/gogf/gf/v2/frame/g"
)

type RequiredEventsReq struct {
	g.Meta `path:"/required_events" tags:"Required Events" method:"get" summary:"Get all events you have processed."`
}

type RequiredEventsRes struct {
	Events []events.EventInfo `json:"events"`
}
