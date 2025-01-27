package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stepbirt/api/todo"
)

type MyContext struct {
	*gin.Context // composition
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c}
}

func (c *MyContext) Bind(v interface{}) error {
	return c.Context.ShouldBindJSON(v)
}
func (c *MyContext) TransactionID() string {
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

func NewGinHandler(handler func(todo.Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Can do both
		// handler(&MyContext{Context: ctx})
		handler(NewMyContext(ctx))
	}
}

type MyRouter struct {
	*gin.Engine
}

func NewMyRouter() *MyRouter {
	r := gin.Default() // have middleware, logging etc
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{
		"http://localhost:3000", // client
	}

	config.AllowHeaders = []string{
		"Origin",
		"Authorization",
		"TransactionID",
	}
	r.Use(cors.New(config))

	return &MyRouter{r}
}

func (r *MyRouter) POST(path string, handler func(todo.Context)) {
	r.Engine.POST(path, NewGinHandler(handler))
}
