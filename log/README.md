# log

> support logrus

## usage

- use interface log

`govendor fetch github.com/s3dteam/go-toolkit/log/^`

- use without interface log

`govendor fetch github.com/s3dteam/go-toolkit/log/logruslogger`

## example

```go
    type temps struct {
		log log.Logger
	}
	options := &logruslogger.Options{}
	options.WithCallerHook = true
	options.Depth = 8

	a := &temps{
		log: logruslogger.GetLoggerWithOptions("a-logrus", options),
	}

	a.log.Debug("")
	a.log.Debug(time.Now())
	a.log.Debug(123, time.Now())
	a.log.Debug("test %v", time.Now().UnixNano())
	a.log.Warn("test %v", time.Now().UnixNano())
	a.log.Info("test %v", time.Now().UnixNano())
	a.log.Printf("test %v", time.Now().UnixNano())
	a.log.Printf("test", time.Now().UnixNano())
	a.log.Error("test %v", time.Now().UnixNano())
```