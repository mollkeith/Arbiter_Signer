package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EventsReq struct {
	g.Meta `path:"/events" tags:"Events" method:"get" summary:"Get all events you have processed."`
}

type EventInfo struct {
	EventID     string
	EventName   string
	EventTime   string
	EventHeight int
}
type EventsRes struct {
	Events []EventInfo `json:"events"`
}
