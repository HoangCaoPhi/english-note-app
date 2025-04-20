package middlewares

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ResponseFormatter struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		if c.IsAborted() {
			return
		}

		// Kiểm tra xem response đã được format chưa
		var originalBody map[string]interface{}
		if blw.body.Len() > 0 {
			if err := json.Unmarshal(blw.body.Bytes(), &originalBody); err != nil {
				return
			}

			// Kiểm tra nếu response đã có format chuẩn thì không wrap nữa
			if _, hasCode := originalBody["code"]; hasCode {
				if _, hasMessage := originalBody["message"]; hasMessage {
					if _, hasData := originalBody["data"]; hasData {
						return
					}
				}
			}
		}

		statusCode := c.Writer.Status()
		var response ResponseFormatter

		switch {
		case statusCode >= 200 && statusCode < 300:
			response = ResponseFormatter{
				Code:    statusCode,
				Message: "success",
				Data:    originalBody,
			}
		case statusCode >= 400 && statusCode < 500:
			errorMsg := ""
			if m, ok := originalBody["error"]; ok {
				errorMsg = m.(string)
			}
			response = ResponseFormatter{
				Code:    statusCode,
				Message: errorMsg,
				Data:    nil,
			}
		default:
			response = ResponseFormatter{
				Code:    statusCode,
				Message: "internal server error",
				Data:    nil,
			}
		}

		c.Header("Content-Type", "application/json")
		c.JSON(statusCode, response)
	}
}
