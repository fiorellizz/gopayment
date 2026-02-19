package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func Init() {
	// Configurando logger para saída em JSON
	Log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}