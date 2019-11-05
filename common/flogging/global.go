package flogging

import (
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
	"strings"
)

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/11/02     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/11/2 3:39 PM
 * @date 2019/11/2 3:39 PM
 * @since 1.0.0
 */

const (
	defaultFormat = "%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}"
	defaultLevel  = zapcore.InfoLevel
)

var Global *Logging
var logger *FabricLogger

func init() {
	logging, err := New(Config{})
	if err != nil {
		panic(err)
	}

	Global = logging
	logger = Global.Logger("flogging")
	grpcLogger := Global.ZapLogger("grpc")
	grpclog.SetLogger(NewGRPCLogger(grpcLogger))
}

// Init initializes logging with the provided config.
func Init(config Config) {
	err := Global.Apply(config)
	if err != nil {
		panic(err)
	}
}

// Reset sets logging to the defaults defined in this package.
//
// Used in tests and in the package init
func Reset() {
	_ = Global.Apply(Config{})
}

// GetLoggerLevel gets the current logging level for the logger with the
// provided name.
func GetLoggerLevel(loggerName string) string {
	return strings.ToUpper(Global.Level(loggerName).String())
}

// MustGetLogger creates a logger with the specified name. If an invalid name
// is provided, the operation will panic.
func MustGetLogger(loggerName string) *FabricLogger {
	return Global.Logger(loggerName)
}

// ActivateSpec is used to activate a logging specification.
func ActivateSpec(spec string) {
	err := Global.ActivateSpec(spec)
	if err != nil {
		panic(err)
	}
}
