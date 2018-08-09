package logruslogger

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func Test_Log(t *testing.T) {
	options := &Options{}
	options.WithCallerHook = true
	options.Depth = 8
	options.Write = true
	tempTestLog := "testLogDir201808091811"
	dir, _ := os.Getwd()

	storeLogDir := filepath.Join(dir, tempTestLog)
	err := os.MkdirAll(storeLogDir, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("creating log file failed: %s", err.Error()))
	}
	options.Path = storeLogDir

	a := GetLoggerWithOptions("a-logrus", options)

	logger := a.GetLogger()
	t.Logf("get logger %v", logger)

	a.Debug("")
	a.Debug(time.Now())
	a.Debug(123, time.Now())
	a.Debug("test %v", time.Now().UnixNano())
	a.Warn("test %v", time.Now().UnixNano())
	a.Info("test %v", time.Now().UnixNano())
	a.Printf("test %v", time.Now().UnixNano())
	a.Printf("test", time.Now().UnixNano())
	a.Error("test %v", time.Now().UnixNano())

	a.Debugln("test", time.Now().UnixNano())
	a.Warnln("test", time.Now().UnixNano())
	a.Infoln("test", time.Now().UnixNano())
	a.Printfln("test", time.Now().UnixNano())
	a.Printfln("test ", time.Now().UnixNano())
	a.Errorln("test", time.Now().UnixNano())

	optionsB := *options
	optionsB.Depth = -1
	optionsB.Formatter = "json"

	b := GetLoggerWithOptions("b-logrus", &optionsB)

	b.Debug("test %v", time.Now().UnixNano())
	b.Warn("test %v", time.Now().UnixNano())
	b.Info("test %v", time.Now().UnixNano())
	b.Printf("test %v", time.Now().UnixNano())
	b.Error("test %v", time.Now().UnixNano())

	b.Debugln("test", time.Now().UnixNano())
	b.Warnln("test", time.Now().UnixNano())
	b.Infoln("test", time.Now().UnixNano())
	b.Printfln("test", time.Now().UnixNano())
	b.Errorln("test", time.Now().UnixNano())

	optionsC := *options
	optionsC.Depth = 19990009900
	c := GetLoggerWithOptions("c-logrus", &optionsC)

	c.Debug("test %v", time.Now().UnixNano())
	c.Warn("test %v", time.Now().UnixNano())

	optionsD := new(Options)

	_ = GetLoggerWithOptions("d-logrus", optionsD)

	optionsE := *options
	optionsE.DisableConsole = true
	_ = GetLoggerWithOptions("e-logrus", &optionsE)

	c.Debug("test %v", time.Now().UnixNano())
	c.Warn("test %v", time.Now().UnixNano())

	os.RemoveAll(options.Path)
}

func Test_PanicLog(t *testing.T) {
	options := &Options{}
	options.WithCallerHook = true
	options.Depth = 8
	options.Write = true
	options.Path = "/logtest"

	a := GetLoggerWithOptions("a-logrus", options)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("fatal log, err: %v\n", err)

			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("fatal log, err: %v\n", err)
				}
			}()
			a.Panicln("fatal test")
		}
	}()

	a.Panic("fatal test")
}
