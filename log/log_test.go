package log_test

import (
	"testing"
	"time"

	"github.com/webbergao1/go-toolkit/log"
	"github.com/webbergao1/go-toolkit/log/logruslogger"
)

func Test_Log(t *testing.T) {
	type temps struct {
		log log.Logger
	}
	options := &logruslogger.Options{}
	options.WithCallerHook = true
	options.Depth = 8

	a := &temps{
		log: logruslogger.GetLoggerWithOptions("a-logrus", options),
	}

	b := &temps{
		log: logruslogger.GetLoggerWithOptions("b-logrus", options),
	}

	a.log.Debug("test %v", time.Now().UnixNano())
	a.log.Warn("test %v", time.Now().UnixNano())
	a.log.Info("test %v", time.Now().UnixNano())
	a.log.Printf("test %v", time.Now().UnixNano())
	a.log.Printf("test", time.Now().UnixNano())
	a.log.Error("test %v", time.Now().UnixNano())

	a.log.Debugln("test", time.Now().UnixNano())
	a.log.Warnln("test", time.Now().UnixNano())
	a.log.Infoln("test", time.Now().UnixNano())
	a.log.Printfln("test", time.Now().UnixNano())
	a.log.Printfln("test ", time.Now().UnixNano())
	a.log.Errorln("test", time.Now().UnixNano())

	b.log.Debug("test %v", time.Now().UnixNano())
	b.log.Warn("test %v", time.Now().UnixNano())
	b.log.Info("test %v", time.Now().UnixNano())
	b.log.Printf("test %v", time.Now().UnixNano())
	b.log.Error("test %v", time.Now().UnixNano())

	b.log.Debugln("test", time.Now().UnixNano())
	b.log.Warnln("test", time.Now().UnixNano())
	b.log.Infoln("test", time.Now().UnixNano())
	b.log.Printfln("test", time.Now().UnixNano())
	b.log.Errorln("test", time.Now().UnixNano())
}
