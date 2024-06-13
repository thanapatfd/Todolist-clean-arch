package middleware

import (
	"log/slog"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// ฟังก์ชันเพื่อแปลง headers จาก map[string][]string เป็น map[string]string
func fiberHeadersToMap(headers map[string][]string) map[string]string {
	// ใช้ lo.MapValues เพื่อแปลงค่าของ headers จาก slice เป็น string
	return lo.MapValues(headers, func(value []string, key string) string {
		return strings.Join(value, ", ") // รวมค่าต่างๆ ใน slice ด้วย ,
	})
}

func Logging(c *fiber.Ctx) error {
	// ใช้ logger ที่กำหนดใน initLogger
	logger := slog.Default()

	start := time.Now()
	err := c.Next()

	// Log ข้อมูลของ request
	logger.Info("Request",
		"Method", c.Method(),
		"Path", c.Path(),
		"Handler", c.Route().Path,
		"Query", c.Queries(),
		"Body", string(c.Body()),
		"Headers", fiberHeadersToMap(c.GetRespHeaders()),
		"Params", c.AllParams(),
		"RemoteIP", c.IP())

	// Log ข้อมูลของ response
	logger.Info("Response",
		"Duration", time.Since(start).Seconds(),
		"Headers", fiberHeadersToMap(c.GetRespHeaders()),
		"Body", string(c.Response().Body()),
		"Status", int64(c.Response().StatusCode()))

	return err
}
