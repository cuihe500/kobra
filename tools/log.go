package tools

import (
	"log/slog"
	"os"
)

var StderrLog *slog.Logger
var StdoutLog *slog.Logger

func init() {
	StdoutLog = slog.New(slog.NewTextHandler(os.Stdout, nil))
	StderrLog = slog.New(slog.NewTextHandler(os.Stderr, nil))
}
