package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(201, gin.H{
		"code":    201,
		"message": "created",
		"data":    data,
	})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"code":    400,
		"message": message,
		"data":    nil,
	})
}

func Unauthorized(c *gin.Context, message string) {
	c.JSON(401, gin.H{
		"code":    401,
		"message": message,
		"data":    nil,
	})
}

func Forbidden(c *gin.Context, message string) {
	c.JSON(403, gin.H{
		"code":    403,
		"message": message,
		"data":    nil,
	})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(404, gin.H{
		"code":    404,
		"message": message,
		"data":    nil,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(500, gin.H{
		"code":    500,
		"message": message,
		"data":    nil,
	})
}
