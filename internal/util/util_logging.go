package util

import (
	"log/slog"
	"os"
)

type Logger struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	From    string `json:"from"`
	Layer   string `json:"layer"`
	TypeLog string `json:"type_log"`
}

func NewLogger(log Logger) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	switch log.TypeLog {
	case "ERROR":
		logger.Error(
			"ERROR",
			"code:", log.Code,
			"message:", log.Message,
			"from:", log.From,
			"layer:", log.Layer,
			"type:", log.TypeLog,
		)
	case "INFO":
		logger.Info(
			"INFO",
			"code:", log.Code,
			"message:", log.Message,
			"from:", log.From,
			"layer:", log.Layer,
			"type:", log.TypeLog,
		)
	case "WARNING":
		logger.Warn(
			"WARNING",
			"code:", log.Code,
			"message:", log.Message,
			"from:", log.From,
			"layer:", log.Layer,
			"type:", log.TypeLog,
		)
	}
}
