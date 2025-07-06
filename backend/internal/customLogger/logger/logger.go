package custlog

import "log/slog"


const (
	LevelTrace = slog.Level(-8)
	LeverFatal = slog.Level(12)
)

type CustomLogger struct {
	logger *slog.Logger
}

func New() *CustomLogger {

}
 
func (l *CustomLogger) 