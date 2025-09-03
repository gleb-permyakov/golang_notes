package logger

import (
	"log"
	"os"
)

type Loga struct {
	infoLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
	warnLogger  *log.Logger
}

func (l *Loga) Info(msg string, args ...interface{}) {
	// printLog(l, "INFO", msg, args)
	if l.infoLogger == nil {
		l.infoLogger = log.New(os.Stderr, "[INFO] ", 0)
	}

	l.infoLogger.Print(msg)
}

func New() *Loga {
	return &Loga{}
}

// func New() *slog.Logger {
// 	env := os.Getenv("ENV")
// 	var levelLogging slog.Level
// 	switch env {
// 	case "dev":
// 		levelLogging = slog.LevelDebug
// 	case "prod":
// 		levelLogging = slog.LevelInfo
// 	}
// 	TextHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
// 		Level:       levelLogging,
// 		ReplaceAttr: customLogger,
// 	})
// 	logger := *slog.New(TextHandler)

// 	return &logger
// }

// func customLogger(groups []string, a slog.Attr) slog.Attr {
// 	// fmt.Println(a.Value)
// 	switch a.Key {
// 	case "time":
// 		return slog.Attr{}
// 	case "level":
// 		typeLog := "[" + a.Value.Any().(slog.Level).String() + "]"
// 		return slog.Attr{Value: slog.StringValue(typeLog)}
// 	case "msg":
// 		return slog.Attr{Value: a.Value}
// 	default:
// 		return a
// 	}
// }
