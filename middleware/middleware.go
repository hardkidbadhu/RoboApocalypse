package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type middleware struct {
	log *logrus.Entry
}

type Middleware interface {
	LoggingHandler(next http.Handler) http.Handler
	RecoverHandler(next http.Handler) http.Handler
}

func NewMiddleware(l *logrus.Entry) *middleware {
	return &middleware{
		l,
	}
}

// RecoverHandler prevent abnormal shutdown while panic
func (l *middleware) RecoverHandler() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				l.log.Printf("panic: %+v", err)
				l.log.Println(string(debug.Stack()))
				http.Error(ctx.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}
