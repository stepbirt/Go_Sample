package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/stepbirt/api/router"
	"github.com/stepbirt/api/store"
	"github.com/stepbirt/api/todo"
)

var (
	buildcommit = "dev"
	buildtime   = time.Now().String()
)

func main() {
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	err = godotenv.Load("local.env")
	if err != nil {
		log.Printf("Please consider environment variables: %s", err)
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DB_CONN")), &gorm.Config{})
	if err != nil {
		panic("fail to connect db")
	}
	db.AutoMigrate(&todo.Todo{})

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017"))
	if err != nil {
		panic("fail to connect db")
	}

	collection := client.Database("myapp").Collection("todos")

	//Gin
	// r := router.NewMyRouter()

	// r.GET("/healthz", func(ctx *gin.Context) {
	// 	ctx.Status(http.StatusOK)
	// })

	// r.GET("/limitz", litmitHandler)

	// r.GET("/tokenz", auth.AccessToken(os.Getenv("SIGNKEY")))

	//fibier
	r := router.NewFiberRouter()
	r.Get("/x", func(ctx *fiber.Ctx) error {
		ctx.Status(http.StatusOK)
		return ctx.SendString("hoho")
	})
	//middleware
	// protectd := r.Group("", auth.Protect([]byte(os.Getenv("SIGNKEY"))))

	// when we use Mysql
	// gormStore := store.NewGormStore(db)
	// handler := todo.NewTodoHandler(gormStore)
	// when we use Mongo
	mongoStore := store.NewMongoDBStore(collection)
	handler := todo.NewTodoHandler(mongoStore)
	r.POST("/todos", handler.NewTask)
	// protectd.GET("/todos", handler.List)
	// protectd.DELETE("/todos/:id", handler.Remove)

	if err := r.Listen(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	// r.Run()
	//graceful shutdown
	// ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// defer stop()
	// s := &http.Server{
	// 	Addr:           ":" + os.Getenv("PORT"),
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// go func() {
	// 	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// <-ctx.Done() // signal from channel both SIGINT, SIGTERM
	// stop()
	// fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	// timeoutCtx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancle()
	// if err := s.Shutdown(timeoutCtx); err != nil {
	// 	fmt.Println(err)
	// }
}

// var limiter = rate.NewLimiter(5, 5)

// func litmitHandler(c *gin.Context) {
// 	if !limiter.Allow() {
// 		c.AbortWithStatus(http.StatusTooManyRequests)
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "ok",
// 	})
// }
