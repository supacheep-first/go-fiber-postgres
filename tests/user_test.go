package tests

import (
	"bytes"
	"encoding/json"
	"fiber-postgre/database"
	"fiber-postgre/models"
	"fiber-postgre/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupApp() *fiber.App {
	godotenv.Load("../.env")
	database.ConnectDB()
	app := fiber.New()
	routes.RegisterUserRoutes(app)
	return app
}

func TestCreateUser(t *testing.T) {
	app := setupApp()

	payload := models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != 201 {
		t.Errorf("expected 201, got %d", resp.StatusCode)
	}
}

func TestGetUsers(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestGetUserNotFound(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, "/users/99999", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}
}
