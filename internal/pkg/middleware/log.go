package middleware

import (
	"bytes"
	"cn.xdmnb/study/miniblog/internal/pkg/log"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.b.Write(b)
	return w.ResponseWriter.Write(b)
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		writer := responseWriter{
			c.Writer,
			bytes.NewBuffer([]byte{}),
		}
		c.Writer = writer
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		clientIP := c.ClientIP()

		//请求体 body
		requestBody := ""
		b, err := c.GetRawData()
		if err != nil {
			requestBody = "failed to get request body"
		} else {
			requestBody = string(b)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		//响应体 body
		responseBody := writer.b.String()
		log.C(c).Infow(reqUri, "latency_time", latencyTime, "client_ip", clientIP, "req_method", reqMethod, "request_body", requestBody, "response_body", responseBody)
	}
}