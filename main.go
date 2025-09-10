package main

import (
	"log/slog"

	"github.com/CarusoVitor/dokuex/cmd"
)

const logLevel = slog.LevelInfo

func main() {
	slog.SetLogLoggerLevel(logLevel)
	cmd.Execute()
}
