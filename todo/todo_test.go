package todo

import (
	"testing"
)

// Old test before apply hexagonal
// func TestCreateTodoNotAllowSleepTask(t *testing.T) {
// 	handler := NewTodoHandler(&gorm.DB{})
// 	w := httptest.NewRecorder()
// 	payload := bytes.NewBufferString(`{"text":"sleep"}`)
// 	req, _ := http.NewRequest("POST", "http://0.0.0.0:8080/todos")
// 	req.Header.Add("TransactionID", "testIDxxx")

// 	c, _ := gin.CreateTestContext(w) //Gin provide create manual context for test
// 	handler.NewTask(c)

// 	want := `{"error":"not allowed"}`

// 	if want != w.Body.String() {
// 		t.Errorf("want %s but get %s\n", want, w.Body.string())
// 	}

// }

func TestCreateTodoNotAllowSleepTask(t *testing.T) {
	handler := NewTodoHandler(&TestDB{})
	c := &TestContext{}
	handler.NewTask(c)

	want := "not allowed"

	if want != c.v["error"] {
		t.Errorf("want %s but get %s\n", want, c.v["error"])
	}

}

// mock dependency
type TestDB struct {
}

func (TestDB) New(*Todo) error {
	return nil
}

type TestContext struct {
	v map[string]interface{}
}

func (TestContext) Bind(v interface{}) error {
	*v.(*Todo) = Todo{
		Title: "sleep",
	}
	return nil
}

func (c *TestContext) JSON(code int, v interface{}) {
	c.v = v.(map[string]interface{})
}

func (TestContext) TransactionID() string {
	return "TestTransactionID"
}

func (TestContext) Username() string {
	return "TestUsername"
}
