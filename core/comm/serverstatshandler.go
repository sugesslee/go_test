package comm

import (
	"context"
	"google.golang.org/grpc/stats"
	"gostudy/metrics"
)

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/11/02     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/11/2 4:44 PM
 * @date 2019/11/2 4:44 PM
 * @since 1.0.0
 */

type ServerStatsHandler struct {
	OpenConnCounter   metrics.Counter
	ClosedConnCounter metrics.Counter
}

func (h *ServerStatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	return ctx
}

func (h *ServerStatsHandler) HandleRPC(ctx context.Context, s stats.RPCStats) {}

func (h *ServerStatsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return ctx
}

func (h *ServerStatsHandler) HandleConn(ctx context.Context, s stats.ConnStats) {
	switch s.(type) {
	case *stats.ConnBegin:
		h.OpenConnCounter.Add(1)
	case *stats.ConnEnd:
		h.ClosedConnCounter.Add(1)
	}
}
