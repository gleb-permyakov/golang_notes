package logger

import (
	"fmt"
	"log"
	"os"
)

var Log Loga

type Level string

const (
	DEV  Level = "dev"
	PROD Level = "prod"
)

type Loga struct {
	envLevel    Level
	infoLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
	warnLogger  *log.Logger
}

func (l *Loga) Info(msg string, args ...interface{}) {
	// Flags of log
	s_Args := validateArgs(args...)

	// log in file
	logInFile(l.infoLogger.Prefix(), msg, s_Args)

	// log in terminal
	l.infoLogger.Print(msg, s_Args)
}

func (l *Loga) Debug(msg string, args ...interface{}) {
	// Flags of log
	s_Args := validateArgs(args...)

	// log in file
	logInFile(l.debugLogger.Prefix(), msg, s_Args)

	// log in terminal
	if l.envLevel == DEV {
		l.debugLogger.Print(msg, s_Args)
	}
}

func (l *Loga) Error(msg string, args ...interface{}) {
	// Flags of log
	s_Args := validateArgs(args...)

	// log in file
	logInFile(l.errorLogger.Prefix(), msg, s_Args)

	// log in terminal
	l.errorLogger.Print(msg, s_Args)
}

func (l *Loga) Warn(msg string, args ...interface{}) {
	// Flags of log
	s_Args := validateArgs(args...)

	// log in file
	logInFile(l.warnLogger.Prefix(), msg, s_Args)

	// log in terminal
	if l.envLevel == DEV {
		l.warnLogger.Print(msg, s_Args)
	}
}

func New() {
	envType := os.Getenv("ENV")
	Log = Loga{
		envLevel:    Level(envType),
		infoLogger:  log.New(os.Stderr, "[INFO] ", 0),
		debugLogger: log.New(os.Stderr, "[DEBUG] ", 0),
		errorLogger: log.New(os.Stderr, "[ERROR] ", 0),
		warnLogger:  log.New(os.Stderr, "[WARN] ", 0),
	}
}

func logInFile(prefix, msg, s_Args string) {
	// path to log_file
	path := os.Getenv("LOG_PATH")

	// open to add log
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(prefix + msg + s_Args + "\n")
	if err != nil {
		panic(err)
	}
}

func validateArgs(args ...interface{}) string {
	var (
		arrArgs []string // all args to str
		s_Args  string   // result string with "key=value" to print
	)

	// Make all args string
	for i := 0; i < len(args); i++ {
		typeAsserter(&arrArgs, &args[i])
	}

	// Make pairs "key=value"
	// for i := 0; i < len(arrArgs)-1; i += 2 {
	// 	key := "" //" " + arrArgs[i] + "=" // if want "key="
	// 	value := (" " + arrArgs[i+1])
	// 	s_Args += (key + value)
	// }
	// ----- OR -------
	// Make string of args without keys
	for i := 0; i < len(arrArgs); i++ {
		s_Args += (" " + arrArgs[i])
	}

	return s_Args
}

func typeAsserter(arrArgs *[]string, a *interface{}) {
	// Asserting different types to string
	switch (*a).(type) {
	case string:
		*arrArgs = append(*arrArgs, (*a).(string))
	case int:
		*arrArgs = append(*arrArgs, fmt.Sprintf("%d", (*a).(int)))
	case int64:
		*arrArgs = append(*arrArgs, fmt.Sprintf("%d", (*a).(int)))
	case float64:
		*arrArgs = append(*arrArgs, fmt.Sprintf("%f", (*a).(float64)))
	case float32:
		*arrArgs = append(*arrArgs, fmt.Sprintf("%f", (*a).(float32)))
	case bool:
		*arrArgs = append(*arrArgs, fmt.Sprintf("%b", (*a).(int)))
	default:
		*arrArgs = append(*arrArgs, "unexpected")
	}
}

// logerr := log.New(os.Stderr, "[ warn  ]", log.Ltime|log.Lshortfile)

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
