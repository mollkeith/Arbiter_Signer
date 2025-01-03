package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SucceedEventsReq struct {
	g.Meta `path:"/succeed_events" tags:"Succeed Events" method:"get" summary:"Get all events succeed."`
}

type SucceedEventInfo struct {
	EventID     string
	EventName   string
	EventTime   string
	EventHeight int
}
type SucceedEventsRes struct {
	Events []EventInfo `json:"events"`
}
