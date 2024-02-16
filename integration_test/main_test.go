package integration_test

import (
	"bytes"
	"log"
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Setup
	payload := []byte(`{"email": "anggriawanrilda88@gmail.com", "password": "sdfsdfsd"}`)
	req, err := http.NewRequest("POST", "/dating-app-test/api/v1/users/", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	log.Println("============", req)

	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(YourCreateUserHandler) // Gantilah dengan fungsi handler yang sesuai

	// // Request
	// handler.ServeHTTP(rr, req)

	// // Assertion
	// assert.Equal(t, http.StatusCreated, rr.Code)
	// // Add more assertions based on the expected response or behavior
}

// func TestUserLogin(t *testing.T) {
// 	// Setup
// 	payload := []byte(`{"username": "testuser", "password": "testpassword"}`)
// 	req, err := http.NewRequest("POST", "/dating-app-test/api/v1/users/login", bytes.NewBuffer(payload))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(YourUserLoginHandler) // Gantilah dengan fungsi handler yang sesuai

// 	// Request
// 	handler.ServeHTTP(rr, req)

// 	// Assertion
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	// Add more assertions based on the expected response or behavior
// }

// func TestGetUserByID(t *testing.T) {
// 	// Setup
// 	req, err := http.NewRequest("GET", "/dating-app-test/api/v1/users/123", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(YourGetUserByIDHandler) // Gantilah dengan fungsi handler yang sesuai

// 	// Request
// 	handler.ServeHTTP(rr, req)

// 	// Assertion
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	// Add more assertions based on the expected response or behavior
// }
