// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package events

import (
	"context"

	"github.com/BeL2Labs/Arbiter_Signer/api/events/v1"
)

type IEventsV1 interface {
	Events(ctx context.Context, req *v1.EventsReq) (res *v1.EventsRes, err error)
	SucceedEvents(ctx context.Context, req *v1.SucceedEventsReq) (res *v1.SucceedEventsRes, err error)
}
