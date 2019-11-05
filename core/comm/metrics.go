package comm

import "gostudy/metrics"

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/11/02     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/11/2 4:43 PM
 * @date 2019/11/2 4:43 PM
 * @since 1.0.0
 */

var (
	openConnCounterOpts = metrics.CounterOpts{
		Namespace: "grpc",
		Subsystem: "comm",
		Name:      "conn_opened",
		Help:      "gRPC connections opened. Open minus closed is the active number of connections.",
	}

	closedConnCounterOpts = metrics.CounterOpts{
		Namespace: "grpc",
		Subsystem: "comm",
		Name:      "conn_closed",
		Help:      "gRPC connections closed. Open minus closed is the active number of connections.",
	}
)

func NewServerStatsHandler(p metrics.Provider) *ServerStatsHandler {
	return &ServerStatsHandler{
		OpenConnCounter:   p.NewCounter(openConnCounterOpts),
		ClosedConnCounter: p.NewCounter(closedConnCounterOpts),
	}
}