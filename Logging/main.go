package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Print("Info message")
	slog.Info("Info message")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// logger.Debug("Debug msg")
	// logger.Info("Info msg")
	// logger.Warn("Warning msg")
	// logger.Error("Error msg")

	// slog.SetDefault(logger)
	// slog.Info("Info msg")
	// logger.Info(
	// 	"incoming request",
	// 	"method", "GET",
	// 	"time_taken_ms", 158,
	// 	"path", "/hello/world?q=search",
	// 	"status", 200,
	// 	"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)")

	// logger.Info(
	// 	"incoming request",
	// 	"method", "GET",
	// 	"time_taken_ms",
	// )

	logger.Info(
		"incoming request",
		slog.String("method", "GET"),
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		slog.Int("status", 200),
		slog.String(
			"user_agent",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	)

}
