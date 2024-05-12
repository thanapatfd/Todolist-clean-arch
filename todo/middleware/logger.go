package middleware

import (
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func fiberHeadersToMap(headers map[string][]string) map[string]string {
	return lo.MapValues(headers, func(value []string, key string) string {
		return strings.Join(value, ", ")
	})
}

func Logging(c *fiber.Ctx) error {

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	start := time.Now()
	err := c.Next()

	logger.Info("Request",
		"Method", c.Method(),
		"Path", c.Path(),
		"Handler", c.Route().Path,
		"Query", c.Queries(),
		"Body", string(c.Body()),
		"Headers", c.GetReqHeaders(),
		"Params", c.AllParams(),
		"RemoteIP", c.IP())

	logger.Info("Request",
		"Duration", time.Since(start).Seconds(),
		"Headers", fiberHeadersToMap(c.GetRespHeaders()),
		"Body", string(c.Response().Body()),
		"Status", int64(c.Response().StatusCode()))

	return err
}
