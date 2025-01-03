package events

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/BeL2Labs/Arbiter_Signer/api/events/v1"
)

func (c *ControllerV1) Events(ctx context.Context, req *v1.EventsReq) (res *v1.EventsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
