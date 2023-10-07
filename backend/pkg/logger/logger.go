package logger

import (
	"os"
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
var once sync.Once
var logCfg = zap.Config{
	Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
	DisableCaller:     false,
	DisableStacktrace: true,
	Sampling: &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 200,
	},
	Encoding:         "json",
	OutputPaths:      []string{"stdout", "my.log"},
	ErrorOutputPaths: []string{"stderr"},
}

type LoggerConfig interface {
	GetAllConfig() []LogPathCfg
}

// various ways for output&encoding
type LogPathCfg struct {
	Encoding string
	Output   string
	Level    string
}

func NewLogPathCfg(encoding, output, level string) LogPathCfg {
	return LogPathCfg{
		Encoding: encoding,
		Output:   output,
		Level:    level,
	}
}

func (l *LogPathCfg) GetAllConfig() []LogPathCfg {
	return []LogPathCfg{*l}
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

// get file descriptor for log file
func getCoreFile(path string) zapcore.WriteSyncer {
	var file *os.File

	switch path {
	case "stdout":
		file = os.Stdout
	case "stderr":
		file = os.Stderr
	default:
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}

		return file
	}

	return file
}

// getLoggerLevel returns zapcore.Level
func getLoggerLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	}

	return zapcore.InfoLevel
}

// GetLogger returns system-named logger.
func GetLogger(system string) *zap.Logger {
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
	return GetLogger(system).Sugar()
}

func GetLoggerSafe(system string) *zap.Logger {
	if rootLogger == nil {
		switch mode {
		case "rich":
			logCfg.EncoderConfig = richEncoder
		default:
			logCfg.EncoderConfig = leanEncoder
		}

		rootLogger, _ = logCfg.Build()
	}

	return rootLogger.Named(system)
}

// LogVerbose dynamically enables/disables log verbosity.
func LogVerbose(enable bool) {
	if enable {
		logCfg.Level.SetLevel(zap.DebugLevel)
	} else {
		logCfg.Level.SetLevel(zap.InfoLevel)
	}
}

func InitLogger(l LoggerConfig) error {
	once.Do(func() {
		lCfg := l.GetAllConfig()
		core := make([]zapcore.Core, 0, len(lCfg))

		for _, k := range lCfg {
			switch k.Encoding {
			case "console":
				consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
				newCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(getCoreFile(k.Output)), getLoggerLevel(k.Level))
				core = append(core, newCore)
			case "json":
				jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
				newCore := zapcore.NewCore(jsonEncoder, zapcore.AddSync(getCoreFile(k.Output)), getLoggerLevel(k.Level))
				core = append(core, newCore)
			}
		}
		combinedCore := zapcore.NewTee(core...)
		rootLogger = zap.New(combinedCore)

		//nolint:gocritic
		defer rootLogger.Sync()
	})

	if rootLogger == nil {
		return ErrCouldNotInit
	}

	return nil
}
