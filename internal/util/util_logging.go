package util

import (
	"log/slog"
	"os"
)

type Layer struct {
	ENTITY_LAYER                               string
	FACTORIES                                  string
	INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION string
	INTERFACE_HANDLERS                         string
	USE_CASES                                  string
}

var LoggerLayers = Layer{
	ENTITY_LAYER: "ENTITY_LAYER",
	FACTORIES:    "FACTORIES",
	INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION: "INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION",
	INTERFACE_HANDLERS:                         "INTERFACE_HANDLERS",
	USE_CASES:                                  "USE_CASES",
}

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
		)
	case "INFO":
		logger.Info(
			"INFO",
			"code:", log.Code,
			"message:", log.Message,
			"from:", log.From,
			"layer:", log.Layer,
		)
	case "WARNING":
		logger.Warn(
			"WARNING",
			"code:", log.Code,
			"message:", log.Message,
			"from:", log.From,
			"layer:", log.Layer,
		)
	}
}
