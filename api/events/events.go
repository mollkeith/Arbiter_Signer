// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package events

import (
	"context"

	"github.com/BeL2Labs/Arbiter_Signer/api/events/v1"
)

type IEventsV1 interface {
	EventLog(ctx context.Context, req *v1.EventLogReq) (res *v1.EventLogRes, err error)
	AllEvents(ctx context.Context, req *v1.AllEventsReq) (res *v1.AllEventsRes, err error)
	FailedEvents(ctx context.Context, req *v1.FailedEventsReq) (res *v1.FailedEventsRes, err error)
	RequiredEvents(ctx context.Context, req *v1.RequiredEventsReq) (res *v1.RequiredEventsRes, err error)
	SucceedEvents(ctx context.Context, req *v1.SucceedEventsReq) (res *v1.SucceedEventsRes, err error)
}
