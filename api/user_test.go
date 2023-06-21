package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/cna-so/todo-api/helpers"
	dto "github.com/cna-so/todo-api/models/DTO"
)

func createUser(router *gin.Engine, db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	userHandlers := initUserApi(db)

	router.POST("/auth", userHandlers.CreateUserAccount)
	userDto := dto.UserSignUpRequestDto{
		Email:     "tesst@mail.com",
		LastName:  "test lastName",
		FirstName: "test firstName",
		Password:  "123456789",
	}
	b, _ := json.Marshal(userDto)
	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	return req, w
}

func TestUserApi_CreateUserAccount(t *testing.T) {
	router := gin.Default()
	db := helpers.SetupDatabase()
	defer helpers.DropTables()

	req, w := createUser(router, db)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
