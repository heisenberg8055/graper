package http_log

import (
	"context"
	"log/slog"
	"os"
)

type Response struct {
	Method     string
	Time_taken string
	Url        string
	Status     int
	Message    string
}

type Err struct {
	Err error
	URL string
}

func LogErr(err Err) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	slog.LogAttrs(context.Background(), slog.LevelError,
		"Info Message",
		slog.String("Error", err.Err.Error()),
		slog.String("URL", err.URL),
	)
}

func LogInfo(response Response) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.LogAttrs(context.Background(), slog.LevelInfo,
		"Info Message",
		slog.String("method", response.Method),
		slog.String("time_taken", response.Time_taken),
		slog.String("path", response.Url),
		slog.Int("status", response.Status),
		slog.String("message", response.Message),
	)
}

func LogError(response Response) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	slog.LogAttrs(context.Background(), slog.LevelError,
		"Info Message",
		slog.String("method", response.Method),
		slog.String("time_taken", response.Time_taken),
		slog.String("path", response.Url),
		slog.Int("status", response.Status),
		slog.String("message", response.Message),
	)
}

func LogWarn(response Response) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.LogAttrs(context.Background(), slog.LevelWarn,
		"Info Message",
		slog.String("method", response.Method),
		slog.String("time_taken", response.Time_taken),
		slog.String("path", response.Url),
		slog.Int("status", response.Status),
		slog.String("message", response.Message),
	)
}
