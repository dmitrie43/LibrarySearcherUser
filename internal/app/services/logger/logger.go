package Logger

import (
	"log/slog"
	"os"

	"github.com/dmitrie43/LibrarySearcherUser/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var log *slog.Logger

func init() {
	cfg := config.MustLoad()
	log = setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func Debug(message string) {
	log.Debug(message)
}
