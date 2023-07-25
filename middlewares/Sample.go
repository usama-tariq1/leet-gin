package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Sample() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")
		// get the value in controller with
		// example := c.MustGet("example").(string)

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
