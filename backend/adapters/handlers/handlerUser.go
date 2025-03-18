package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"orderkuota/dto"
	"orderkuota/middleware"
	ports "orderkuota/ports/service"
	"orderkuota/utils"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "failed", "Method not allowed")
		return
	}

	var userDTO dto.UserRegister
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "failed", "Invalid request body")
		return
	}

	if !utils.IsValidEmail(userDTO.Email) {
		utils.ErrorResponse(w, http.StatusBadRequest, "failed", "Invalid email format")
		return
	}

	if len(userDTO.Password) < 6 {
		utils.ErrorResponse(w, http.StatusBadRequest, "failed", "Password must be at least 6 characters")
		return
	}

	user := userDTO.ToDomain()
	createdUser, err := uh.Service.RegisterUser(user)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	response := dto.FromDomain(createdUser)
	utils.ResponseJSON(w, response, http.StatusCreated, "success", "User registered")
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "failed", "Method not allowed")
		return
	}

	var loginDTO dto.UserLogin
	err := json.NewDecoder(r.Body).Decode(&loginDTO)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "failed", "Invalid body request")
		return
	}

	if !utils.IsValidEmail(loginDTO.Email) {
		utils.ErrorResponse(w, http.StatusBadRequest, "failed", "Invalid email format")
		return
	}

	user := loginDTO.ToDomain()
	loggedInUser, token, err := uh.Service.Login(user.Email, user.Password)
	if err != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, "failed", "Invalid credentials")
		return
	}

	userResponse := dto.FromDomain(loggedInUser)

	loginResponse := dto.LoginResponse{
		User:        userResponse,
		AccessToken: token,
	}

	utils.ResponseJSON(w, loginResponse, http.StatusOK, "success", "Login successful")
}

func (uh *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "failed", "Method not allowed")
		return
	}

	users, err := uh.Service.GetAllUsers()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	utils.ResponseJSON(w, users, http.StatusOK, "success", "Users retrieved successfully")
}

func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "failed", "Method not allowed")
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || id == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "failed", "User ID is required")
		return
	}

	user, err := uh.Service.GetUserByID(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "failed", "User not found")
		return
	}

	utils.ResponseJSON(w, user, http.StatusOK, "success", "User retrieved successfully")
}

func (uh *UserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPut {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "failed", "Method not allowed")
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || id == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "failed", "User ID is required")
		return
	}

	var updateDTO dto.UserUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&updateDTO)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "failed", "Invalid request body")
		return
	}

	existingUser, err := uh.Service.GetUserByID(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "failed", "User not found")
		return
	}

	if updateDTO.Email != nil {
		if !utils.IsValidEmail(*updateDTO.Email) {
			utils.ErrorResponse(w, http.StatusBadRequest, "failed", "Invalid email format")
			return
		}
		existingUser.Email = *updateDTO.Email
	}
	if updateDTO.Name != nil {
		existingUser.Name = *updateDTO.Name
	}
	if updateDTO.Password != nil {
		existingUser.Password = *updateDTO.Password
	}

	updatedUser, err := uh.Service.UpdateUserByID(id, existingUser)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}

	utils.ResponseJSON(w, updatedUser, http.StatusOK, "success", "User updated successfully")
}

func (uh *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "failed", "Method not allowed")
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || id == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "failed", "User ID is required")
		return
	}

	loggedInUserID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || loggedInUserID == "" {
		utils.ErrorResponse(w, http.StatusUnauthorized, "failed", "Unauthorized")
		return
	}

	err := uh.Service.DeleteUserByID(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed", "Failed to delete user")
		log.Println("DeleteUserByID error:", err)
		return
	}

	if id == loggedInUserID {
		response := map[string]interface{}{
			"status":  "success",
			"message": "User deleted successfully. Please logout.",
		}
		utils.ResponseJSON(w, response, http.StatusOK, "success", "User deleted successfully")
		return
	}

	utils.ResponseJSON(w, nil, http.StatusOK, "success", "User deleted successfully")
}
