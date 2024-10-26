package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"learn-go/models"
	"learn-go/repositories"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users := h.userRepo.List()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("User not found")
		return
	}
	user := h.userRepo.Get(int64(id))

	if user == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("User not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	if name == "" || email == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Name and email are required")
		return
	}

	newUser := models.User{
		Name:  name,
		Email: email,
	}

	newId := h.userRepo.Create(newUser)
	user := h.userRepo.Get(newId)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("User not found")
		return
	}
	user := h.userRepo.Get(int64(id))

	if user == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("User not found")
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}

	h.userRepo.Update(int64(id), *user)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("User not found")
		return
	}
	user := h.userRepo.Get(int64(id))

	if user == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("User not found")
		return
	}

	h.userRepo.Delete(int64(id))

	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
}
