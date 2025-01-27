package todo

import "github.com/gin-gonic/gin"

type MyContext struct {
	*gin.Context // composition
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c}
}

func (c *MyContext) Bind(v interface{}) error {
	return c.Context.ShouldBindJSON(v)
}
func (c *MyContext) TransactionId() string {
	return c.Request.Header.Get("TransactionID")
}
func (c *MyContext) Username() string {

	if username, ok := c.Get("username"); ok {
		if s, ok := username.(string); ok {
			return s
		}
	}
	return ""
}
func (c *MyContext) JSON(code int, v interface{}) {
	c.Context.JSON(code, v)
}

func NewGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Can do both
		// handler(&MyContext{Context: ctx})
		handler(NewMyContext(ctx))
	}
}
