package todo

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
