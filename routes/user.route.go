package routes

import (
	"database/sql"
	"learn-go/handlers"
	"learn-go/repositories"
	"net/http"
)

func UserRoutes(router *http.ServeMux, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	router.HandleFunc("GET /users", userHandler.List)
	router.HandleFunc("GET /users/{userId}", userHandler.Get)
	router.HandleFunc("POST /users", userHandler.Create)
	router.HandleFunc("PATCH /users/{userId}", userHandler.Update)
	router.HandleFunc("DELETE /users/{userId}", userHandler.Delete)
}
