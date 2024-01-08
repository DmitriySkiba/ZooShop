package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func TestRegistration(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("name", "John Doe")
	body.Add("login", "johndoe")
	body.Add("password", "password123")
	req, err := http.NewRequest("POST", "/registration", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	// Вызов функции для тестирования
	db.Registration(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Success" {
		t.Errorf("Expected 'Success' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestRegistration1(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("name", "John Doe")
	body.Add("login", "johndoe")
	//body.Add("password", "password123")
	req, err := http.NewRequest("POST", "/registration", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	// Вызов функции для тестирования
	db.Registration(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Error executing query" {
		t.Errorf("Expected 'Error executing query' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestName(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "johndoe")
	req, err := http.NewRequest("POST", "/name", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	// Вызов функции для тестирования
	db.Name(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "John Doe" {
		t.Errorf("Expected 'John Doe' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestName2(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "john")
	req, err := http.NewRequest("POST", "/name", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	// Вызов функции для тестирования
	db.Name(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Error executing query" {
		t.Errorf("Expected 'Error executing query' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestAuthorisation(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "johndoe")
	body.Add("password", "password123")
	req, err := http.NewRequest("POST", "/authorisation", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	// Вызов функции для тестирования
	db.Authorisation(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Valid login and password" {
		t.Errorf("Expected 'Valid login and password' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestAuthorisation2(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "346")
	req, err := http.NewRequest("POST", "/authorisation", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	// Вызов функции для тестирования
	db.Authorisation(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Invalid login or password" {
		t.Errorf("Expected 'Invalid login or password' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestAuthorisation3(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "346")
	body.Add("password", "456")
	req, err := http.NewRequest("POST", "/authorisation", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	// Вызов функции для тестирования
	db.Authorisation(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Invalid login or password" {
		t.Errorf("Expected 'Invalid login or password' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestSale(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "johndoe")
	req, err := http.NewRequest("POST", "/sale", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	//  Вызов функции для тестирования
	db.Sale(rr, req)

	res, _ := strconv.ParseFloat(rr.Body.String(), 64)
	// Проверка ожидаемого результата
	if res != 0 {
		t.Errorf("Expected '0' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestBought(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "johndoe")
	body.Add("sum", "110000")
	req, err := http.NewRequest("POST", "/bought", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	// Вызов функции для тестирования
	db.Bought(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Success" {
		t.Errorf("Expected 'Success' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestBought2(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "john")
	body.Add("sum", "110000")
	req, err := http.NewRequest("POST", "/bought", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	// Вызов функции для тестирования
	db.Bought(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Error" {
		t.Errorf("Expected 'Error' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestSale2(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "johndoe")
	req, err := http.NewRequest("POST", "/sale", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	//  Вызов функции для тестирования
	db.Sale(rr, req)

	res, _ := strconv.ParseFloat(rr.Body.String(), 64)
	// Проверка ожидаемого результата
	if res != 0.1 {
		t.Errorf("Expected '0.1' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestSale3(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "john")
	req, err := http.NewRequest("POST", "/sale", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	//  Вызов функции для тестирования
	db.Sale(rr, req)

	res := rr.Body.String()
	// Проверка ожидаемого результата
	if res != "Error executing query" {
		t.Errorf("Expected 'Error executing query' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestChangePw(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "johndoe")
	body.Add("password", "123")
	req, err := http.NewRequest("POST", "/change_pw", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	//  Вызов функции для тестирования
	db.ChangePw(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Success" {
		t.Errorf("Expected 'Success' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestChangePw2(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "john")
	body.Add("password", "123")
	req, err := http.NewRequest("POST", "/change_pw", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	//  Вызов функции для тестирования
	db.ChangePw(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Error" {
		t.Errorf("Expected 'Error' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestChangeName(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "johndoe")
	body.Add("name", "doejohn")
	req, err := http.NewRequest("POST", "/change_name", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	//  Вызов функции для тестирования
	db.ChangeName(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Success" {
		t.Errorf("Expected 'Success' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestChangeName2(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "john")
	body.Add("name", "doejohn")
	req, err := http.NewRequest("POST", "/change_name", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	//  Вызов функции для тестирования
	db.ChangeName(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Error" {
		t.Errorf("Expected 'Error' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestCheckStatus(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "johndoe")
	req, err := http.NewRequest("POST", "/check_status", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	//  Вызов функции для тестирования
	db.CheckStatus(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "0" {
		t.Errorf("Expected '0' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestCheckStatus2(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "john")
	req, err := http.NewRequest("POST", "/check_status", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	//  Вызов функции для тестирования
	db.CheckStatus(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Error executing query" {
		t.Errorf("Expected 'Error executing query' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestDelete(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "johndoe")
	req, err := http.NewRequest("POST", "/delete", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	//  Вызов функции для тестирования
	db.Delete(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Success" {
		t.Errorf("Expected 'Success' but got %s", rr.Body.String())
	}
	db.Close()
}

func TestDelete2(t *testing.T) {
	// Создание экземпляра базы данных для тестов
	db := NewDatabase()

	// Создание тестового HTTP запрос
	body := url.Values{}
	body.Add("login", "john")
	req, err := http.NewRequest("POST", "/delete", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Создание тестового HTTP response writer
	rr := httptest.NewRecorder()

	//  Вызов функции для тестирования
	db.Delete(rr, req)

	// Проверка ожидаемого результата
	if rr.Body.String() != "Error" {
		t.Errorf("Expected 'Error' but got %s", rr.Body.String())
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
		{name: "Registration", path: "/registration", expectedStatus: http.StatusOK},
		{name: "Authorisation", path: "/authorisation", expectedStatus: http.StatusOK},
		{name: "Sale", path: "/sale", expectedStatus: http.StatusOK},
		{name: "Bought", path: "/bought", expectedStatus: http.StatusOK},
		{name: "Delete", path: "/delete", expectedStatus: http.StatusOK},
		{name: "ChangePw", path: "/change_pw", expectedStatus: http.StatusOK},
		{name: "ChangeName", path: "/change_name", expectedStatus: http.StatusOK},
		{name: "Name", path: "/name", expectedStatus: http.StatusOK},
		{name: "CheckStatus", path: "/check_status", expectedStatus: http.StatusOK},
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