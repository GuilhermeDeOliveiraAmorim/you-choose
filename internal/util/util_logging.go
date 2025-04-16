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
	CONFIGURATION                              string
	MIDDLEWARES                                string
}

type TypeLog struct {
	ERROR   string
	INFO    string
	WARNING string
}

var LoggerLayers = Layer{
	ENTITY_LAYER: "ENTITY_LAYER",
	FACTORIES:    "FACTORIES",
	INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION: "INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION",
	INTERFACE_HANDLERS:                         "INTERFACE_HANDLERS",
	USE_CASES:                                  "USE_CASES",
	CONFIGURATION:                              "CONFIGURATION",
	MIDDLEWARES:                                "MIDDLEWARES",
}

var LoggerTypes = TypeLog{
	ERROR:   "ERROR",
	INFO:    "INFO",
	WARNING: "WARNING",
}

type Logger struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	From    string `json:"from"`
	Layer   string `json:"layer"`
	TypeLog string `json:"type_log"`
}

var logger *slog.Logger

func InitLogger() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func NewLogger(log Logger) {
	if logger == nil {
		InitLogger()
	}

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
