package middleware

import (
	"fmt"
	"log/slog"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func HandlePanic() fiber.Handler {
	return recover.New(recover.Config{
		Next:             nil,
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e any) {
			slog.Error(
				"Panic Recovered",
				"err", fmt.Errorf("%v", e),
				"stack", debug.Stack(),
			)
		},
	})
}
