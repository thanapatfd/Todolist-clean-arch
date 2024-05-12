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

// Logging middleware logs the request path, method, and request processing time.
func Logging(c *fiber.Ctx) error {

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	start := time.Now()
	err := c.Next()

	// สร้าง payload ของ request โดยใช้ map
	requestPayload := map[string]interface{}{
		"Method":   c.Method(),
		"Path":     c.Path(),
		"Handler":  c.Route().Path,
		"Query":    c.Queries(),
		"Body":     string(c.Body()),
		"Headers":  fiberHeadersToMap(c.GetReqHeaders()),
		"Params":   c.AllParams(),
		"RemoteIP": c.IP(),
	}

	responsePayload := map[string]interface{}{
		"Duration": time.Since(start).Seconds(),
		"Headers":  fiberHeadersToMap(c.GetRespHeaders()),
		"Body":     string(c.Response().Body()),
		"Status":   int64(c.Response().StatusCode()),
	}
	// บันทึก payload ของ request ด้วย slog.Info
	println(c.Response().StatusCode())
	if c.Response().StatusCode() == 200 {
		slog.Info("Request received", requestPayload)
		slog.Info("Request completed", responsePayload)

	}
	if c.Response().StatusCode() == 400 {
		slog.Warn("Request received", requestPayload)
		slog.Warn("Request completed", responsePayload)

	}
	if c.Response().StatusCode() == 500 {
		slog.Error("Request received", requestPayload)
		slog.Error("Request completed", responsePayload)

	}

	return err
}
