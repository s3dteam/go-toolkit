package logruslogger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const (
	defaultLevel       = "level"
	defaultLogFileName = "all.logs"
)

var logMap map[string]*LogrusLogger
var getLogMutex sync.Mutex

func defaultOptions() *Options {
	return &Options{
		Level:          defaultLevel,
		Depth:          8,
		WithCallerHook: true,
		Formatter:      "text",
		DisableConsole: false,
		Write:          false,
		Path:           os.TempDir(),
		FileName:       defaultLogFileName,
		MaxAge:         time.Duration(24) * time.Hour,
		RotationTime:   time.Duration(7*24) * time.Hour,
	}
}

// Options logger options config
type Options struct {
	Level          string
	Depth          int
	WithCallerHook bool
	Formatter      string // only support json and text

	DisableConsole bool
	Write          bool
	Path           string
	FilePrefix     string
	FileName       string

	MaxAge       time.Duration
	RotationTime time.Duration
}

// GetLoggerWithOptions with options config
func GetLoggerWithOptions(logName string, options *Options) *LogrusLogger {
	getLogMutex.Lock()
	defer getLogMutex.Unlock()

	if options == nil {
		options = defaultOptions()
	}

	if logMap == nil {
		logMap = make(map[string]*LogrusLogger)
	}
	curLog, ok := logMap[logName]

	if ok {
		return curLog
	}

	log := logrus.New()

	// get logLevel
	level := options.Level
	if level == "" {
		level = defaultLevel
	}
	logLevel := GetLogLevel(level)
	logDir := options.Path
	logFileName := options.FileName
	printLog := !options.DisableConsole
	depth := options.Depth
	maxAge := options.MaxAge
	rotationTime := options.RotationTime
	withCallerHook := options.WithCallerHook
	defaultLogFilePrex := options.FilePrefix

	log.SetLevel(logLevel)

	if options.Write {
		storeLogDir := filepath.Join(logDir)

		err := os.MkdirAll(storeLogDir, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("creating log file failed: %s", err.Error()))
		}

		path := filepath.Join(storeLogDir, logFileName)
		writer, err := rotatelogs.New(
			path+".%Y%m%d%H%M%S",
			rotatelogs.WithClock(rotatelogs.Local),
			rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour),
		)
		if err != nil {
			panic(fmt.Sprintf("rotatelogs log failed: %s", err.Error()))
		}

		var formatter logrus.Formatter

		formatter = &logrus.TextFormatter{}
		if options.Formatter == "json" {
			formatter = &logrus.JSONFormatter{}
		}

		log.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.DebugLevel: writer,
				logrus.InfoLevel:  writer,
				logrus.WarnLevel:  writer,
				logrus.ErrorLevel: writer,
				logrus.FatalLevel: writer,
			},
			formatter,
		))

		pathMap := lfshook.PathMap{
			logrus.DebugLevel: fmt.Sprintf("%s/%sdebug.log", storeLogDir, defaultLogFilePrex),
			logrus.InfoLevel:  fmt.Sprintf("%s/%sinfo.log", storeLogDir, defaultLogFilePrex),
			logrus.WarnLevel:  fmt.Sprintf("%s/%swarn.log", storeLogDir, defaultLogFilePrex),
			logrus.ErrorLevel: fmt.Sprintf("%s/%serror.log", storeLogDir, defaultLogFilePrex),
			logrus.FatalLevel: fmt.Sprintf("%s/%sfatal.log", storeLogDir, defaultLogFilePrex),
		}
		log.AddHook(lfshook.NewHook(
			pathMap,
			formatter,
		))
	} else {
		if printLog {
			log.Out = os.Stdout
		}
	}

	if withCallerHook {
		log.AddHook(&CallerHook{depth: depth, module: logName}) // add caller hook to print caller's file and line number
	}
	curLog = &LogrusLogger{
		log: log,
	}
	logMap[logName] = curLog
	fmt.Printf("register logger %v, current loggers: %v\n", logName, logMap)
	return curLog
}
