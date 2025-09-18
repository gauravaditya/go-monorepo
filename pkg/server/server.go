package server

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const defaultReadBufSize = 8192

type Config struct {
	Fiber fiber.Config
	// Otel  otel.Config // Uncomment if you want to allow custom Otel config
}

func New(serviceName string, config ...Config) *fiber.App {
	app := fiber.New(fiber.Config{AppName: serviceName, ReadBufferSize: defaultReadBufSize})
	// tracingMiddleware := otel.NewMiddleware(serviceName)

	if len(config) > 0 {
		cfg := config[0].Fiber
		if cfg.AppName == "" {
			cfg.AppName = serviceName
		}

		if cfg.ReadBufferSize == 0 {
			cfg.ReadBufferSize = defaultReadBufSize
		}

		app = fiber.New(cfg)
		// tracingMiddleware = otel.NewMiddleware(serviceName, config[0].Otel)
	}

	app.Use(
		// tracingMiddleware.OtelFiberHandler(),
		// tracingMiddleware.TraceIdResponseHeaderHandler(), //nolint:lll // should be after tracing is configured by Open Telemetry Fiber middleware
		logger.New(loggerConfig(serviceName)), // should be after trace-id is added in response header
		// tracingMiddleware.SessionIdHandler(),
		// middleware.HandlePanic(),
		// middleware.AddCustomerIPToUserContext(),
	)

	return app
}

func loggerConfig(serviceName string) logger.Config {
	fields := map[string]string{
		"time":           "${time}",
		"level":          "INFO",
		"msg":            "request completed",
		"status":         "${status}",
		"method":         "${method}",
		"path":           "${path}",
		"correlation-id": "${reqHeader:correlation-Id}",
		"service-name":   serviceName,
		"trace_id":       "${respHeader:trace-id}",
	}

	format, err := json.Marshal(fields)
	if err != nil {
		panic(fmt.Errorf("unable to create logger config: %w", err))
	}

	return logger.Config{
		TimeFormat: time.RFC3339Nano,
		Format:     fmt.Sprintf("%s\n", string(format)),
	}
}
