package helpers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/luisfelipegodoi/go-superhero-lib/src/context/general/utils"
	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/lrhook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Logger struct {
	Severity     int
	PackageName  string
	Method       string
	ParentMethod string
	Message      string
	TraceID      string
}

type StackTrace struct {
	Severity      string   `json:"level"`
	TraceID       string   `json:"trace_id"`
	Message       string   `json:"message"`
	PackageNames  []string `json:"package_names"`
	ParentMethods []string `json:"parent_methods"`
	Methods       []string `json:"methods"`
	Messages      []string `json:"messages"`
	Time          string   `json:""`
}

var (
	rootLogger *logrus.Logger
)

const (
	// LevelDebug - Level logging = debug
	LevelDebug = iota
	// LevelInfo - Level logging = info
	LevelInfo = iota
	// LevelWarn - Level logging = warn
	LevelWarn = iota
	// LevelError - Level logging = error
	LevelError = iota
	// LevelFatal - Level logging = fatal
	LevelFatal = iota
	// LevelPanic - Level logging = panic
	LevelPanic = iota
)

var levels = map[int]string{
	LevelDebug: "Debug",
	LevelInfo:  "Info",
	LevelWarn:  "Warn",
	LevelError: "Error",
	LevelFatal: "Fatal",
	LevelPanic: "Panic",
}

var wasConfigure = false

// ConfigureLogger configure log handler
func ConfigureLogger() {
	configureLogrus()
	configureSlack()
	wasConfigure = true
}

func New(severity int, packageName, parentMethod, method, message, traceID string) Logger {
	return Logger{
		Severity:     severity,
		PackageName:  packageName,
		ParentMethod: parentMethod,
		Method:       method,
		Message:      message,
		TraceID:      traceID,
	}
}

func createMessage(stackTrace StackTrace) (string, error) {
	msg, err := json.Marshal(stackTrace)

	if err != nil {
		LogIt(LevelError, "logger", "createMessage", "json.Marshal", err.Error())
		return "", err
	}

	return fmt.Sprintf("%s ```%s```", stackTrace.Message, string(msg)), err
}

// FirstMessage - SendLog
func FirstMessage(loggers []Logger) string {
	if len(loggers) > 0 {
		return loggers[0].Message
	}

	return ""
}

// VerifyAndSendLogByLevel - Verify the level of the first message and send the log if the first log level be lower and equal than
// the parameter level, true will be returned too.
func VerifyAndSendLogByLevel(level int, sendLog bool, loggers []Logger) bool {
	if len(loggers) > 0 && loggers[0].Severity <= level {
		if sendLog {
			SendLog(loggers[0].PackageName, loggers[0].ParentMethod, loggers[0].Method, loggers[0].Message, loggers[0].TraceID, loggers)
		}

		return true
	}

	return false
}

//SendLog send log
func SendLog(packageName, parentMethod, method, message, traceID string, loggers []Logger) {
	stackTrace := StackTrace{}

	if len(loggers) > 0 {

		if traceID == "" && loggers[0].TraceID != "" {
			traceID = loggers[0].TraceID
		}

		if message == "" {
			message = loggers[0].Message
		}

		stackTrace.Severity = levels[loggers[0].Severity]
		stackTrace.TraceID = traceID
		stackTrace.Message = message
		for _, logger := range loggers {
			stackTrace.PackageNames = append(stackTrace.PackageNames, logger.PackageName)
			stackTrace.ParentMethods = append(stackTrace.ParentMethods, logger.ParentMethod)
			stackTrace.Methods = append(stackTrace.Methods, logger.Method)
			stackTrace.Messages = append(stackTrace.Messages, logger.Message)
		}
	} else {
		LogIt(LevelInfo, packageName, parentMethod, method, message, traceID)
		return
	}

	msg, err := createMessage(stackTrace)

	if err != nil {
		msg = message
	}

	LogIt(loggers[0].Severity, packageName, parentMethod, method, msg, traceID)
}

// LogIt logs on to logrus
// params[0] -> TraceID
// params[1] -> []Logger ou Bool, caso Bool e true ele irá "logar" a mensagem, caso
func LogIt(severity int, packageName, parentMethod, method, message string, params ...interface{}) []Logger {
	if !wasConfigure {
		ConfigureLogger()
	}

	var logger Logger
	var loggers []Logger
	var loggerObjectIndex = -1
	var sendLog = false
	var converted = false
	var traceID = ""

	if len(params) > 0 {
		traceID, converted = params[0].(string)
		if len(params) == 2 {
			loggerObjectIndex = 1
		} else {
			if !converted {
				// significa que no primeiro parametro não foi passado o TraceID e sim o objeto de Logger ou a flag.
				loggerObjectIndex = 0
			}
		}
	}

	if loggerObjectIndex >= 0 {
		loggers, converted = params[loggerObjectIndex].([]Logger)

		if converted {
			loggers = append(loggers, New(severity, packageName, parentMethod, method, message, traceID))
			return loggers
		} else {
			logger, converted = params[loggerObjectIndex].(Logger)
			if converted {
				loggers = append(loggers, logger)
				loggers = append(loggers, New(severity, packageName, parentMethod, method, message, traceID))
				return loggers
			} else {
				sendLog, converted = params[loggerObjectIndex].(bool)
				if converted && !sendLog {
					loggers = append(loggers, New(severity, packageName, parentMethod, method, message, traceID))
					return loggers
				}
			}
		}
	}

	logrusEntry := rootLogger.WithFields(withFields(packageName, parentMethod, method, traceID))

	switch severity {
	case LevelDebug:
		logrusEntry.Debug(message)
	case LevelInfo:
		logrusEntry.Info(message)
	case LevelWarn:
		logrusEntry.Warn(message)
	case LevelError:
		logrusEntry.Error(message)
	case LevelFatal:
		logrusEntry.Fatal(message)
	case LevelPanic:
		logrusEntry.Panic(message)
	default:
		logrusEntry.Debug(message)
	}

	loggers = append(loggers, New(severity, packageName, parentMethod, method, message, traceID))
	return loggers

}

func withFields(packageName, parentMethod, method, traceID string) logrus.Fields {
	return logrus.Fields{
		"app":           viper.Get("application.name"),
		"env":           viper.Get("application.env"),
		"kind":          viper.Get("application.kind"),
		"team":          viper.Get("application.team"),
		"type":          "json",
		"package":       packageName,
		"parent_method": parentMethod,
		"method":        method,
		"trace": logrus.Fields{
			"id": traceID,
		},
	}
}

func configureLogrus() {
	rootLogger = logrus.New()
	rootLogger.Formatter = &logrus.JSONFormatter{TimestampFormat: time.RFC3339}
	rootLogger.SetLevel(getLogLevel("slack.loglevel"))
}

func configureSlack() {

	config := lrhook.Config{
		MinLevel: getLogLevel("slack.loglevel"),
		Message: chat.Message{
			Channel:   utils.InterfaceToString(viper.Get("slack.channel")),
			IconEmoji: utils.InterfaceToString(viper.Get("slack.notification-icon")),
			Username:  utils.InterfaceToString(viper.Get("slack.username")),
			Markdown:  false,
		},
	}
	hook := lrhook.New(config, utils.InterfaceToString(viper.Get("slack.webhook-url")))
	rootLogger.AddHook(hook)
}

func getLogLevel(envVariable string) logrus.Level {
	switch viper.Get(envVariable) {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "FATAL":
		return logrus.FatalLevel
	case "PANIC":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}
