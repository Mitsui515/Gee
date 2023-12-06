package geeweb

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start Time
		t := time.Now()
		// Process Request
		c.Next()
		// Calculate Resolution Time
		log.Printf("[%d] %s in %v\n", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
