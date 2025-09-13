package main

import (
	"log/slog"

	"github.com/CarusoVitor/dokuex/cmd"
)

const logLevel = slog.LevelDebug

func main() {
	slog.SetLogLoggerLevel(logLevel)
	cmd.Execute()
}
