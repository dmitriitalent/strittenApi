package loggerService

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(message string, args ...any)
	Error(message string, args ...any)
	Debug(message string, args ...any)
}

type LoggerService struct {
	sugar *zap.SugaredLogger
}

func NewLoggerService() *LoggerService {
	level := zapcore.DebugLevel

encoderCfg := zapcore.EncoderConfig{
    TimeKey:        "ts",
    LevelKey:       "level",
    NameKey:        "logger",
    CallerKey:      "caller",
    MessageKey:     "msg",
    StacktraceKey:  "stacktrace",
    LineEnding:     zapcore.DefaultLineEnding,
    EncodeLevel:    zapcore.CapitalLevelEncoder,
    EncodeTime:     zapcore.EpochNanosTimeEncoder, 
    EncodeDuration: zapcore.SecondsDurationEncoder,
    EncodeCaller:   zapcore.ShortCallerEncoder,
}

	consoleWriter := zapcore.Lock(os.Stdout)
    if err := os.MkdirAll("logs", 0755); err != nil {
        panic("cannot create logs directory: " + err.Error())
    }
    
    file, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic("cannot open log file: " + err.Error())
    }

    fileWriter := zapcore.AddSync(file)

    encoder := zapcore.NewJSONEncoder(encoderCfg)

    core := zapcore.NewTee(
        zapcore.NewCore(encoder, consoleWriter, level),
        zapcore.NewCore(encoder, fileWriter, level),
    )

    logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
    sugar := logger.Sugar()

    return &LoggerService{sugar: sugar}
}

func (l *LoggerService) Info(message string, args ...any) {
	l.sugar.Infof(message, args...)
}

func (l *LoggerService) Error(message string, args ...any) {
	l.sugar.Errorf(message, args...)
}

func (l *LoggerService) Debug(message string, args ...any) {
	l.sugar.Debugf(message, args...)
}

func (l *LoggerService) Sync() error {
	err := l.sugar.Sync()
	if err != nil && !strings.Contains(err.Error(), "The handle is invalid") {
		return err
	}
	return nil
}
