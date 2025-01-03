// Copyright (c) 2025 The bel2 developers
package v1

import (
	"github.com/BeL2Labs/Arbiter_Signer/utility/events"

	"github.com/gogf/gf/v2/frame/g"
)

type FailedEventsReq struct {
	g.Meta `path:"/failed_events" tags:"All Events" method:"get" summary:"Get all events failed."`
}

type FailedEventsRes struct {
	Events []events.EventInfo `json:"events"`
}
