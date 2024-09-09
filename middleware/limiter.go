package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	visitors map[int64]*Visitor
	mu       sync.Mutex
	rate     time.Duration
	burst    int
}

type Visitor struct {
	limiter  *time.Ticker
	lastSeen time.Time
}

// NewRateLimiter creates a new rate limiter with the given rate and burst.
func NewRateLimiter(rate time.Duration, burst int) *RateLimiter {
	return &RateLimiter{
		visitors: make(map[int64]*Visitor),
		rate:     rate,
		burst:    burst,
	}
}

// getVisitor returns the visitor for the given userID. If the visitor does not exist, it creates a new one.
func (rl *RateLimiter) getVisitor(userID int64) *Visitor {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[userID]
	if !exists {
		limiter := time.NewTicker(rl.rate)
		rl.visitors[userID] = &Visitor{limiter: limiter, lastSeen: time.Now()}
		return rl.visitors[userID]
	}

	v.lastSeen = time.Now()
	return v
}

// cleanupVisitors removes visitors that have not been seen for more than 3 minutes.
func (rl *RateLimiter) cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		rl.mu.Lock()
		for id, v := range rl.visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(rl.visitors, id)
			}
		}
		rl.mu.Unlock()
	}
}

// Limit returns a middleware that limits the request rate for each user.
func (rl *RateLimiter) Limit() gin.HandlerFunc {
	go rl.cleanupVisitors()

	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
			c.Abort()
			return
		}

		visitor := rl.getVisitor(userID.(int64))

		select {
		case <-visitor.limiter.C:
			c.Next()
		default:
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
		}
	}
}
