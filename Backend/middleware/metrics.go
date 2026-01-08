package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	mu        sync.Mutex
	reqCount  int
	lastReset = time.Now()
)

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		reqCount++
		mu.Unlock()

		c.Next()
	}
}

func GetQPS() float64 {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(lastReset).Seconds()
	if elapsed == 0 {
		return 0
	}

	qps := float64(reqCount) / elapsed

	reqCount = 0
	lastReset = now

	return qps
}
