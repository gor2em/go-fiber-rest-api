package user

import (
	"fmt"
	"go-fiber-rest-api/pkg/config"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetHandler(t *testing.T){
	database, err := config.ConnectDB()

	assert.Nil(t, err)
	repo := NewRepository(database)
	service := NewService(repo)
	handler := NewHandler(service)

	app := fiber.New()

	app.Get("/users/:id", handler.Get)
	id, err := repo.Create(Model{Name:"test", Email: "test@mail.com"})
	assert.Nil(t, err)

	req:= httptest.NewRequest("GET", fmt.Sprintf("/users/%d",id), nil)
	resp, err := app.Test(req)
	assert.Nil(t,err)
	assert.Equal(t, 200, resp.StatusCode) //beklenen-alınan

}