package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*	1. Подключиться к БД
	2. Создать товар
	3. Отредактировать товар
	4. Выбрать товар
	5. Добавить товар в корзину
	6. Просмотр корзины
	7. Удалить из корзины
	8. Добавить в корзину
	9. Купить товар
	10. История покупок
	11. Удалить товар
	12. Лента
*/
// Подключение к базе=================================================================================================================================
func TestConnect(t *testing.T) {
	// Вызов функции для тестирования
	db := NewDatabase()

	// Проверка, что объект базы данных не равен nil
	if db == nil {
		t.Errorf("Expected non-nil database, but got nil")
	}

	// Проверка, что соединение с базой данных открыто
	if db.db.Ping() != nil {
		t.Errorf("Expected database connection to be open")
	}

	// Закрытие соединения с базой данных
	db.Close()

	// Проверка, что соединение с базой данных закрыто
	if db.db.Ping() == nil {
		t.Errorf("Expected database connection to be closed")
	}
}

// Новый товар ==========================================================================================================================================================
func TestRegisterNewGoods1(t *testing.T) {
	// Создание тестовых данных товара
	db := NewDatabase()
	newGoods := Goods{
		Name:        "test_name",
		Price:       100.0,
		Description: "test_description",
		Photo:       "test_photo",
		Category:    "test_category",
		Subcategory: "test_subcategory",
		Brand:       "test_brand",
		Status:      "admin",
	}

	jsonData, err := json.Marshal(newGoods)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.register_new_goods(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "New goods added successfully"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestRegisterNewGoods2(t *testing.T) {
	db := NewDatabase()
	// Создание тестовых данных товара
	newGoods := Goods{
		Name:        "test_name",
		Price:       100.0,
		Description: "test_description",
		Photo:       "test_photo",
		Category:    "test_category",
		Subcategory: "test_subcategory",
		Brand:       "test_brand",
		Status:      "user",
	}

	jsonData, err := json.Marshal(newGoods)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.register_new_goods(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "User have not got enough laws"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestRegisterNewGoods3(t *testing.T) {
	db := NewDatabase()
	// Создание тестовых данных товара
	newGoods := Goods{
		Name:        "",
		Price:       100.0,
		Description: "test_description",
		Photo:       "test_photo",
		Category:    "test_category",
		Subcategory: "test_subcategory",
		Brand:       "test_brand",
		Status:      "admin",
	}

	jsonData, err := json.Marshal(newGoods)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.register_new_goods(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "Some fields are null"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestRegisterNewGoods4(t *testing.T) {
	db := NewDatabase()
	// Создание тестовых данных товара
	newGoods := Goods{
		Name:        "Корм",
		Price:       0,
		Description: "test_description",
		Photo:       "test_photo",
		Category:    "test_category",
		Subcategory: "test_subcategory",
		Brand:       "test_brand",
		Status:      "admin",
	}

	jsonData, err := json.Marshal(newGoods)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.register_new_goods(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "Some fields are null"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestRegisterNewGoods5(t *testing.T) {
	db := NewDatabase()
	// Создание тестовых данных товара
	newGoods := Goods{
		Name:        "Корм",
		Price:       100.0,
		Description: "test_description",
		Photo:       "test_photo",
		Category:    "test_category",
		Subcategory: "test_subcategory",
		Brand:       "",
		Status:      "admin",
	}

	jsonData, err := json.Marshal(newGoods)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.register_new_goods(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "Some fields are null"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestRegisterNewGoods6(t *testing.T) {
	db := NewDatabase()
	// Создание тестовых данных товара
	newGoods := Goods{
		Name:        "Корм",
		Price:       100.0,
		Description: "test_description",
		Photo:       "test_photo",
		Category:    "test_category",
		Subcategory: "test_subcategory",
		Brand:       "test_brand",
		Status:      "ERR",
	}

	jsonData, err := json.Marshal(newGoods)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.register_new_goods(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "Unknown status"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}

// Удалить товар ========================================================================================================================================================
func TestDeleteGood1(t *testing.T) {
	db := NewDatabase()
	// Подготовка тестовых данных
	good := Delete_good{
		Status: "admin",
		Vcode:  123,
	}
	jsonData, err := json.Marshal(good)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/delete", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.delete_good(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "Good was deleted"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestDeleteGood2(t *testing.T) {
	db := NewDatabase()
	// Подготовка тестовых данных
	good := Delete_good{
		Status: "user",
		Vcode:  123,
	}
	jsonData, err := json.Marshal(good)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/delete", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.delete_good(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "User have not enough laws"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestDeleteGood3(t *testing.T) {
	db := NewDatabase()
	// Подготовка тестовых данных
	good := Delete_good{
		Status: "ERR",
		Vcode:  123,
	}
	jsonData, err := json.Marshal(good)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/delete", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.delete_good(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "Unknown status of person"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}

// Редактирование товара ============================================================================================================================================
func TestUpdateGood1(t *testing.T) {
	db := NewDatabase()
	//fmt.Print("!")
	// Создаем тестовый товар
	good := Goods{
		Status:      "admin",
		Name:        "Test Good",
		Price:       100.0,
		Description: "Test Description",
		Photo:       "Test Photo",
		Category:    "Test Category",
		Subcategory: "Test Subcategory",
		Brand:       "Test Brand",
		Vcode:       4452327765, //существующй vcode
	}

	// Преобразуем товар в JSON
	jsonData, err := json.Marshal(good)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/update", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.update_good(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "Goods was updated successfully"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestUpdateGood2(t *testing.T) {
	db := NewDatabase()
	// Создаем тестовый товар
	good := Goods{
		Status:      "user",
		Name:        "Test Good",
		Price:       100.0,
		Description: "Test Description",
		Photo:       "Test Photo",
		Category:    "Test Category",
		Subcategory: "Test Subcategory",
		Brand:       "Test Brand",
		Vcode:       4452327765,
	}

	// Преобразуем товар в JSON
	jsonData, err := json.Marshal(good)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/update", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.update_good(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "User have not enough laws"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}

func TestUpdateGood3(t *testing.T) {
	db := NewDatabase()
	// Создаем тестовый товар
	good := Goods{
		Status:      "admin",
		Name:        "",
		Price:       100.0,
		Description: "Test Description",
		Photo:       "Test Photo",
		Category:    "Test Category",
		Subcategory: "Test Subcategory",
		Brand:       "Test Brand",
		Vcode:       123,
	}

	// Преобразуем товар в JSON
	jsonData, err := json.Marshal(good)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/update", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.update_good(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "Some fields filled incorrect"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestUpdateGood4(t *testing.T) {
	db := NewDatabase()
	// Создаем тестовый товар
	good := Goods{
		Status:      "admin",
		Name:        "Test Good",
		Price:       100.0,
		Description: "Test Description",
		Photo:       "Test Photo",
		Category:    "Test Category",
		Subcategory: "Test Subcategory",
		Brand:       "Test Brand",
		Vcode:       123, //несуществующий vcode
	}

	// Преобразуем товар в JSON
	jsonData, err := json.Marshal(good)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/update", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.update_good(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	// Проверка сообщения
	expected := "No rows were updated"
	if rr.Body.String() != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}

// Выборка товара ================================================================================================================================================================================
func TestSelectGood1(t *testing.T) {
	db := NewDatabase()
	// Подготовка тестовых данных
	good := Vcode_good{
		Vcode: 4452327765,
	}
	jsonData, err := json.Marshal(good)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/select", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.select_good(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	// Проверка содержимого ответа
	expectedJson := `"{\"name\" : \"Test Good\", \"price\" : 100, \"description\" : \"Test Description\", \"photo\" : \"Test Photo\", \"category\" : \"Test Category\", \"subcategory\" : \"Test Subcategory\", \"brand\" : \"Test Brand\", \"vcode\" : 4452327765}"
`
	if rr.Body.String() != expectedJson {
		t.Errorf("Expected json '%s', got '%s'", expectedJson, rr.Body.String())
	}
	db.Close()
}

// Добавить в корзину ==============================================================================================================================================================================
func TestAdd2basket1(t *testing.T) {
	db := NewDatabase()
	// Подготовка тестовых данных
	add_good := BasketAction{
		IdPerson: 287,
		Vcode:    4452327765,
	}
	jsonData, err := json.Marshal(add_good)
	if err != nil {
		t.Fatal(err)
	}
	// Создание тестового запроса
	req, _ := http.NewRequest("POST", "/add2basket", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.add2basket(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	// Проверка содержимого ответа
	expected := "New goods added successfully into basket"
	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}

// Удалить из корзины ============================================================================================================================================================================
func TestDeleteInBasket1(t *testing.T) {
	db := NewDatabase()
	// Подготовка тестовых данных
	delete_good := BasketAction{
		IdPerson: 287,
		Vcode:    4452327765,
	}
	jsonData, err := json.Marshal(delete_good)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req, _ := http.NewRequest("POST", "/deleteInBasket", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.delete_in_basket(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	// Проверка содержимого ответа
	expected := "Good was deleted from basket successfully"
	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}

// Для просмотра корзины================================================================================================================================
func TestAdd2basket2(t *testing.T) {
	db := NewDatabase()
	// Подготовка тестовых данных
	add_good := BasketAction{
		IdPerson: 287,
		Vcode:    4452327765,
	}
	jsonData, err := json.Marshal(add_good)
	if err != nil {
		t.Fatal(err)
	}
	// Создание тестового запроса
	req, _ := http.NewRequest("POST", "/add2basket", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.add2basket(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	// Проверка содержимого ответа
	expected := "New goods added successfully into basket"
	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}

// Корзина пользователя ===========================================================================================================================================================
func TestPersonalBasket1(t *testing.T) {
	db := NewDatabase()
	// Подготовка тестовых данных
	view := Id_person{
		Id: 287,
	}
	jsonData, err := json.Marshal(view)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/viewBasket", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.personal_basket(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	// Проверка содержимого ответа
	expectedJson := `"[{\"name\" : \"Test Good\", \"price\" : 100, \"photo\" : \"Test Photo\", \"vcode\" : 4452327765}]"
`
	if rr.Body.String() != expectedJson {
		t.Errorf("Expected json '%s', got '%s'", expectedJson, rr.Body.String())
	}
	db.Close()
}

// Удаление перед пустой корзиной
func TestDeleteInBasket2(t *testing.T) {
	db := NewDatabase()
	// Подготовка тестовых данных
	delete_good := BasketAction{
		IdPerson: 287,
		Vcode:    4452327765,
	}
	jsonData, err := json.Marshal(delete_good)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req, _ := http.NewRequest("POST", "/deleteInBasket", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.delete_in_basket(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	// Проверка содержимого ответа
	expected := "Good was deleted from basket successfully"
	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestPersonalBasket2(t *testing.T) {
	db := NewDatabase()
	// Подготовка тестовых данных
	view := Id_person{
		Id: 287,
	}
	jsonData, err := json.Marshal(view)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/viewBasket", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.personal_basket(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	// Проверка содержимого ответа
	expected := `"Basket is empty"
`
	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.db.Close()
}

// Купить товар ====================================================================================================================================================================
func TestAdd2basket3(t *testing.T) {
	db := NewDatabase()
	// Подготовка тестовых данных
	add_good := BasketAction{
		IdPerson: 287,
		Vcode:    4452327765,
	}
	jsonData, err := json.Marshal(add_good)
	if err != nil {
		t.Fatal(err)
	}
	// Создание тестового запроса
	req, _ := http.NewRequest("POST", "/add2basket", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.add2basket(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	// Проверка содержимого ответа
	expected := "New goods added successfully into basket"
	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}

func TestBuy(t *testing.T) {
	db := NewDatabase()

	bought_good := BasketAction{
		Vcode:    4452327765,
		IdPerson: 287,
	}
	jsonData, err := json.Marshal(bought_good)
	if err != nil {
		t.Fatal(err)
	}
	// Создание тестового запроса
	req, _ := http.NewRequest("POST", "/buy", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.buy(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	// Проверка содержимого ответа
	expected := "Good was bought"
	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}

// История покупок ===================================================================================================================================================================
func TestStoryPurchased(t *testing.T) {
	db := NewDatabase()

	// Prepare test data
	view := Id_person{
		Id: 287,
	}
	jsonData, err := json.Marshal(view)
	if err != nil {
		t.Fatal(err)
	}

	// Create test request
	req := httptest.NewRequest("POST", "/storyBought", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Create test response
	rr := httptest.NewRecorder()

	// Call function directly
	db.story_purchased(rr, req)

	// Check response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	// Check response content
	expectedJson := `"[{\"name\" : \"Test Good\", \"price\" : 100, \"photo\" : \"Test Photo\", \"vcode\" : 4452327765}]"
`
	if rr.Body.String() != expectedJson {
		t.Errorf("Expected json '%s', got '%s'", expectedJson, rr.Body.String())
	}
	db.Close() // Close the database connection
}

// Лента общая ====================================================================================================================================================================
func TestLenta1(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "date",
		Column:  "",
		Meaning: "",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	db.Close()
}
func TestLenta2(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "DESC",
		Column:  "",
		Meaning: "",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	expected := "It can be type of sort in common input"

	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestLenta3(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "date",
		Column:  "category",
		Meaning: "собаки",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	db.Close()
}
func TestLenta4(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "date",
		Column:  "category",
		Meaning: "",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	expected := "Incorrect struct of sorting json"

	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestLenta5(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "ASC",
		Column:  "category",
		Meaning: "собаки",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	db.Close()
}
func TestLenta6(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "DESC",
		Column:  "category",
		Meaning: "собаки",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	db.Close()
}
func TestLenta7(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "ERR",
		Column:  "category",
		Meaning: "собаки",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	expected := "Incorrect type of sort"

	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestLenta8(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "DESC",
		Column:  "category",
		Meaning: "",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	expected := "Incorrect struct of sorting json"

	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestLenta9(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "ASC",
		Column:  "name",
		Meaning: "",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	db.Close()
}
func TestLenta10(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "DESC",
		Column:  "name",
		Meaning: "",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	db.Close()
}
func TestLenta11(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "ERR",
		Column:  "name",
		Meaning: "",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	expected := "Invalid type of sort"

	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestLenta12(t *testing.T) {
	db := NewDatabase()
	object := Lenta{
		Sort:    "ASC",
		Column:  "name",
		Meaning: "ERR",
	}
	jsonData, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}

	// Создание тестового запроса
	req := httptest.NewRequest("POST", "/allAndSort", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	// Создание тестового ответа
	rr := httptest.NewRecorder()

	// Вызов функции напрямую
	db.lenta(rr, req)

	// Проверка ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
	expected := "Incorrect, there is no MEANING"

	if rr.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rr.Body.String())
	}
	db.Close()
}
func TestSetupServer(t *testing.T) {
	db := NewDatabase()
	defer db.Close()

	server := SetupServer(db)

	ts := httptest.NewServer(server.Handler)
	defer ts.Close()

	tests := []struct {
		name           string
		path           string
		expectedStatus int
	}{
		{name: "register_new_goods", path: "/register", expectedStatus: http.StatusOK},
		{name: "delete_good", path: "/delete", expectedStatus: http.StatusOK},
		{name: "update_good", path: "/update", expectedStatus: http.StatusOK},
		{name: "select_good", path: "/select", expectedStatus: http.StatusOK},
		{name: "add2basket", path: "/addTobasket", expectedStatus: http.StatusOK},
		{name: "delete_in_basket", path: "/deleteInBasket", expectedStatus: http.StatusOK},
		{name: "personal_basket", path: "/viewBasket", expectedStatus: http.StatusOK},
		{name: "buy", path: "/buy", expectedStatus: http.StatusOK},
		{name: "story_purchased", path: "/storyBought", expectedStatus: http.StatusOK},
		{name: "lenta", path: "/allAndSort", expectedStatus: http.StatusOK},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := http.Get(ts.URL + test.path)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != test.expectedStatus {
				t.Errorf("Expected status code %d but got %d", test.expectedStatus, resp.StatusCode)
			}
		})
	}
}
