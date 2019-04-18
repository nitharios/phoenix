package middleware

import (
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
)

type (
	// RequestHeader represents structure of incoming request headers
	RequestHeader struct {
		mutex sync.RWMutex
	}
)

// NewRequestHeaders for passing request headers
func NewRequestHeaders() *RequestHeader {
	return &RequestHeader{}
}

// Process is middleware that handles processing request headers
func (rh *RequestHeader) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

		h := c.Request().Header

		token := h.Get("token")
		ct := h.Get("Content-Type")
		fmt.Println(token)
		fmt.Println(ct)
		// rh.mutex.Lock()
		// defer rh.mutex.Unlock()

		if token != "" {
			return nil
		}

		return echo.ErrForbidden
	}
}
