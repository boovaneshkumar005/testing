package functional

import (
	"bytes"
	"encoding/json"
	"gotesting/config"
	"gotesting/handler"
	"gotesting/repository"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullUserFlow(t *testing.T) {
	config.InitDB()
	defer config.CloseDB()

	// 1. Add user via HTTP
	newUser := repository.User{ID: 600, Name: "FunctionalDBUser"}
	body, _ := json.Marshal(newUser)
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler.AddUserHandler(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// 2. Retrieve user via HTTP
	reqGet := httptest.NewRequest(http.MethodGet, "/user?id=600", nil)
	wGet := httptest.NewRecorder()
	handler.GetUserHandler(wGet, reqGet)

	assert.Equal(t, http.StatusOK, wGet.Code)
	var fetchedUser repository.User
	json.NewDecoder(wGet.Body).Decode(&fetchedUser)
	assert.Equal(t, "FunctionalDBUser", fetchedUser.Name)
}

// Benchmark for the full user flow
func BenchmarkFullUserFlow(b *testing.B) {
	config.InitDB()
	defer config.CloseDB()

	for i := 0; i < b.N; i++ {
		// Add user
		newUser := repository.User{ID: i, Name: "BenchmarkUser"}
		body, _ := json.Marshal(newUser)
		req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(body))
		w := httptest.NewRecorder()
		handler.AddUserHandler(w, req)

		// Get user
		reqGet := httptest.NewRequest(http.MethodGet, "/user?id="+strconv.Itoa(i), nil)
		wGet := httptest.NewRecorder()
		handler.GetUserHandler(wGet, reqGet)
	}
}
