package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func New(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderXXSSProtection, "1; mode=block")
		c.Set(fiber.HeaderXContentTypeOptions, "nosniff")
		c.Set(fiber.HeaderXDownloadOptions, "noopen")
		c.Set(fiber.HeaderStrictTransportSecurity, "max-age=5184000")
		c.Set(fiber.HeaderXFrameOptions, "SAMEORIGIN")
		c.Set(fiber.HeaderXDNSPrefetchControl, "off")

		return c.Next()
	})
	app.Use(limiter.New(limiter.Config{
		Max: 2000000,
	}))
	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(pprof.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	format := `{"event_time":"${time}", "source_ip":"${ip}", "status":${status}, "latency":"${latency}", "method":"${method}", "path":"${path}", "request_id":"${locals:requestid}",`
	format = format + `"bytes_sent":${bytesSent}, "bytes_received":${bytesReceived}, "user_agent":"${ua}"}` + "\n"
	app.Use(logger.New(logger.Config{
		Format: format,
	}))
}
