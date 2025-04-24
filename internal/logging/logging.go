package logging

import (
	"context"
	"log/slog"
	"os"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/lmittmann/tint"
)

type Layer struct {
	ENTITY                                     string
	FACTORIES                                  string
	INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION string
	INTERFACE_HANDLERS                         string
	USECASES                                   string
	CONFIGURATION                              string
	MIDDLEWARES                                string
}

type TypeLog struct {
	ERROR   string
	INFO    string
	WARNING string
}

var LoggerLayers = Layer{
	ENTITY:    "ENTITY",
	FACTORIES: "FACTORIES",
	INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION: "INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION",
	INTERFACE_HANDLERS:                         "INTERFACE_HANDLERS",
	USECASES:                                   "USECASES",
	CONFIGURATION:                              "CONFIGURATION",
	MIDDLEWARES:                                "MIDDLEWARES",
}

var LoggerTypes = TypeLog{
	ERROR:   "ERROR",
	INFO:    "INFO",
	WARNING: "WARNING",
}

type Logger struct {
	Context  context.Context             `json:"context"`
	Code     int                         `json:"code"`
	Message  string                      `json:"message"`
	From     string                      `json:"from"`
	Layer    string                      `json:"layer"`
	TypeLog  string                      `json:"type_log"`
	Error    error                       `json:"error"`
	Problems []exceptions.ProblemDetails `json:"problems"`
}

var logger *slog.Logger

func InitLogger() {
	logger = slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug, // ou slog.LevelInfo
		TimeFormat: "2006-01-02 15:04:05",
	}))
}

func NewLogger(log Logger) {
	if logger == nil {
		InitLogger()
	}

	switch log.TypeLog {
	case "ERROR":
		logger.ErrorContext(
			log.Context,
			"ERROR",
			"code:", log.Code,
			"message:", log.Message,
			"from:", log.From,
			"layer:", log.Layer,
			"error:", log.Error,
			"problems:", log.Problems,
		)
	case "INFO":
		logger.InfoContext(
			log.Context,
			"INFO",
			"code:", log.Code,
			"message:", log.Message,
			"from:", log.From,
			"layer:", log.Layer,
			"error:", log.Error,
			"problems:", log.Problems,
		)
	case "WARNING":
		logger.WarnContext(
			log.Context,
			"WARNING",
			"code:", log.Code,
			"message:", log.Message,
			"from:", log.From,
			"layer:", log.Layer,
			"error:", log.Error,
			"problems:", log.Problems,
		)
	}
}
