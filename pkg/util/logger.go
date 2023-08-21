package util

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Mode defines logging output configuration.
// Supported modes: lean/rich.
// Set it w/ build-pipeline to modify logging format.
var mode = "rich"

// Root logger not intended for direct usage,
// it rather defines common configuration for
// all subsystem-specific child loggers.
var rootLogger *zap.Logger

var logCfg = zap.Config{
	Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
	DisableCaller:     false,
	DisableStacktrace: true,
	Sampling: &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 200,
	},
	Encoding:         "console",
	OutputPaths:      []string{"stdout", "my.log"},
	ErrorOutputPaths: []string{"stderr"},
}

// fully fledged logging output
var richEncoder = zap.NewDevelopmentEncoderConfig()

// no event time in output, default for syslog-based environments
var leanEncoder = zapcore.EncoderConfig{
	TimeKey:        zapcore.OmitKey,
	LevelKey:       "L",
	NameKey:        "N",
	CallerKey:      "C",
	FunctionKey:    zapcore.OmitKey,
	MessageKey:     "M",
	StacktraceKey:  zapcore.OmitKey,
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

// subsystem loggers
var sLoggers = make(map[string]*zap.Logger)
var sLoggerMx sync.Mutex

// GetLogger returns system-named logger.
func Getlogger(system string) *zap.Logger {
	sLoggerMx.Lock()
	defer sLoggerMx.Unlock()

	logger, exist := sLoggers[system]

	if !exist {
		logger = rootLogger.Named(system)
		sLoggers[system] = logger
	}

	return logger
}

// GetSugaredLogger returns system-named sugared logger.
func GetSugaredLogger(system string) *zap.SugaredLogger {
	return Getlogger(system).Sugar()
}

// LogVerbose dynamically enables/disables log verbosity.
func LogVerbose(enable bool) {
	if enable {
		logCfg.Level.SetLevel(zap.DebugLevel)
	} else {
		logCfg.Level.SetLevel(zap.InfoLevel)
	}
}

//nolint:gochecknoinits
func init() {
	switch mode {
	case "rich":
		logCfg.EncoderConfig = richEncoder
	default:
		logCfg.EncoderConfig = leanEncoder
	}

	rootLogger, _ = logCfg.Build()
}
