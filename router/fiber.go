package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/stepbirt/api/todo"
)

type FiberRouter struct {
	*fiber.App
}

func NewFiberRouter() *FiberRouter {
	r := fiber.New()

	r.Use(cors.New())
	r.Use(logger.New())

	return &FiberRouter{r}
}

func (f *FiberRouter) POST(path string, handler func(todo.Context)) {
	f.App.Post(path, func(c *fiber.Ctx) error {
		handler(NewFiberContext(c))
		return nil
	})
}

func (f *FiberRouter) GET(path string, handler func(todo.Context)) {
	f.App.Get(path, func(c *fiber.Ctx) error {
		handler(NewFiberContext(c))
		return nil
	})
}

type FiberContext struct {
	*fiber.Ctx // composition
}

func NewFiberContext(c *fiber.Ctx) *FiberContext {
	return &FiberContext{Ctx: c}
}

func (c *FiberContext) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}
func (c *FiberContext) TransactionID() string {
	return string(c.Ctx.Request().Header.Peek("TransactionID"))
}
func (c *FiberContext) Username() string {
	return c.Ctx.Get("username")
}
func (c *FiberContext) JSON(code int, v interface{}) {
	c.Ctx.Status(code).JSON(v)
}
