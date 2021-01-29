package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct{
	c *gin.Context
}


func NewResponseWriter(c *gin.Context)*Response{
	return &Response{
		c:c,
	}
}



func (r *Response) Json(data interface{})error{
	r.c.JSON(http.StatusOK,gin.H{
		"data":data,
	})
	return nil
}
