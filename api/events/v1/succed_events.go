package v1

import (
	"github.com/BeL2Labs/Arbiter_Signer/utility/events"

	"github.com/gogf/gf/v2/frame/g"
)

type SucceedEventsReq struct {
	g.Meta `path:"/succeed_events" tags:"Succeed Events" method:"get" summary:"Get all events succeed."`
}

type SucceedEventsRes struct {
	Events []events.EventInfo `json:"events"`
}
