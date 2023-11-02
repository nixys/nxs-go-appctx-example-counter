package ctx

import (
	"fmt"
	"os"
	"time"

	appctx "github.com/nixys/nxs-go-appctx/v3"
	"github.com/sirupsen/logrus"
)

// Ctx defines application custom context
type Ctx struct {
	Conf            confOpts
	CounterInterval time.Duration
	Log             *logrus.Logger
}

func AppCtxInit() (any, error) {

	c := &Ctx{}

	args, err := ArgsRead()
	if err != nil {
		return nil, err
	}

	conf, err := confRead(args.ConfigPath)
	if err != nil {
		// Write to temp logger
		tmpLogError("ctx init: %s", err.Error())
		return nil, err
	}

	c.Log, err = logInit(conf.LogFile, conf.LogLevel)
	if err != nil {
		tmpLogError("ctx init: %s", err.Error())
		return nil, err
	}

	// Set application context
	c.Conf = conf

	if args.CounterInterval != nil {
		c.CounterInterval, err = time.ParseDuration((*args.CounterInterval))
		if err != nil {
			c.Log.Errorf("ctx init: parse counter interval: %s", err.Error())
			return nil, err
		}
	} else {
		c.CounterInterval, err = time.ParseDuration(conf.CounterInterval)
		if err != nil {
			c.Log.Errorf("ctx init: parse counter interval: %s", err.Error())
			return nil, err
		}
	}

	return c, nil
}

func tmpLogError(format string, args ...interface{}) {
	l, _ := appctx.DefaultLogInit(os.Stderr, logrus.InfoLevel, &logrus.JSONFormatter{})
	l.Errorf(format, args...)
}

func logInit(file, level string) (*logrus.Logger, error) {

	var (
		f   *os.File
		err error
	)

	switch file {
	case "stdout":
		f = os.Stdout
	case "stderr":
		f = os.Stderr
	default:
		f, err = os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
		if err != nil {
			return nil, fmt.Errorf("log init: %w", err)
		}
	}

	// Validate log level
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, fmt.Errorf("log init: %w", err)
	}

	return appctx.DefaultLogInit(f, l, &logrus.JSONFormatter{})
}
