package main

import "log/slog"
import "os"

func main() {
	cfg := config{
		address: ":8080",
		db:      dbConfig{},
	}
	api := application{
		config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.Info("Starting API server...")
	slog.SetDefault(logger)
	handler := api.mount()
	err := api.run(handler)
	if err != nil {
		slog.Error("There was an error: ", "error", err) // There's no need to have this followed by an "os.Exit(1)", as log.Fatal is just log.Print followed by an os.Exit(1).
	}
}
