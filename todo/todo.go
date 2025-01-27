package todo

import (
	"log"
	"net/http"
	"time"
)

type Todo struct {
	Title     string `json:"text"` //frontend will send text not title
	ID        uint   `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt DeletedAt `gorm:"index"`
}

// Naming Table
func (Todo) TableName() string {
	return "todos"
}

type TodoHandler struct {
	store storer
}

type storer interface {
	New(*Todo) error
}

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{})
	TransactionID() string
	Username() string
}

func NewTodoHandler(store storer) *TodoHandler {
	return &TodoHandler{store: store}
}

func (t *TodoHandler) NewTask(c Context) {
	var todo Todo
	// if err := c.ShouldBindJSON(&todo); err != nil {
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if todo.Title == "sleep" {
		// transactionID := c.Request.Header.Get("TransactionID")
		transactionID := c.TransactionID()
		// username, _ := c.Get("username")
		username := c.Username()
		log.Println(transactionID, username, "not allowed")
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "not allowed",
		})
		return
	}

	err := t.store.New(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"ID": todo.ID,
	})
}

// func (t *TodoHandler) List(c *gin.Context) {
// 	var todos []Todo
// 	r := t.db.Find(&todos)
// 	if err := r.Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, todos)
// }

// func (t *TodoHandler) Remove(c *gin.Context) {
// 	idParam := c.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	r := t.db.Delete(&Todo{}, id)

// 	if err := r.Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, map[string]interface{}{
// 		"status": "success",
// 	})
// }
