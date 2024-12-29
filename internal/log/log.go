package http_log

import (
	"context"
	"log/slog"
	"os"
	"sync"

	"github.com/lmittmann/tint"
)

type Response struct {
	Method  string
	Url     string
	Status  int
	Message string
}

type Err struct {
	Err error
	URL string
}

type Map struct {
	Mp map[string]bool
	mu sync.Mutex
}

func NewMap() *Map {
	return &Map{Mp: make(map[string]bool)}
}

func (m *Map) Set(key string, value bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Mp[key] = value
}

func (m *Map) Get(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.Mp[key]
	return ok
}

func LogErr(message string, err Err) {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	slog.LogAttrs(context.Background(), slog.LevelError,
		message,
		slog.String("Error", err.Err.Error()),
		slog.String("URL", err.URL),
	)
}

func LogInfo(response Response) {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	slog.LogAttrs(context.Background(), slog.LevelInfo,
		"Info Message",
		slog.String("method", response.Method),
		slog.String("path", response.Url),
		slog.Int("status", response.Status),
		slog.String("message", response.Message),
	)
}

func LogError(response Response) {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	slog.LogAttrs(context.Background(), slog.LevelError,
		"Info Message",
		slog.String("method", response.Method),
		slog.String("path", response.Url),
		slog.Int("status", response.Status),
		slog.String("message", response.Message),
	)
}

func LogWarn(response Response) {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	slog.LogAttrs(context.Background(), slog.LevelWarn,
		"Info Message",
		slog.String("method", response.Method),
		slog.String("path", response.Url),
		slog.Int("status", response.Status),
		slog.String("message", response.Message),
	)
}
