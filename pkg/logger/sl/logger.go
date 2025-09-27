package sl

import (
	"io"
	stdLog "log"
	"log/slog"
	"os"
)

type Type int

const (
	DefaultType Type = iota
	TextType
	JSONType
)

type LogLevel struct {
	level string
}

func SetDefaultLogger(loggerType Type, serviceName string, opts *slog.HandlerOptions) {
	var loggerHandler slog.Handler
	switch loggerType {
	case TextType:
		loggerHandler = slog.NewTextHandler(os.Stdout, opts)
	case JSONType, DefaultType:
		loggerHandler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(loggerHandler).With(slog.String("serviceName", serviceName))
	slog.SetDefault(logger)
}

type HandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

type Handler struct {
	slog.Handler
	l     *stdLog.Logger
	attrs []slog.Attr
}

func (opts HandlerOptions) NewHandler(out io.Writer) *Handler {
	return &Handler{Handler: slog.NewJSONHandler(out, opts.SlogOpts), l: stdLog.New(out, "", 0)}
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func Set(level string, debug bool) {
	switch level {
	case "local", "storage", "client", "frontend":
		opts := HandlerOptions{SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug}}
		slog.SetDefault(slog.New(opts.NewHandler(os.Stdout)))
	default:
		loggerType := JSONType
		loggerOpts := &slog.HandlerOptions{}

		if debug {
			loggerType = TextType
			loggerOpts = &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}
		}

		SetDefaultLogger(loggerType, "easycart", loggerOpts)
	}
}
