package user

import (
	"fmt"
	"net/http"

	"github.com/kasariks/simple_messenger_api/config"
	"github.com/kasariks/simple_messenger_api/services/auth"
	"github.com/kasariks/simple_messenger_api/types"
	"github.com/kasariks/simple_messenger_api/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /register", h.HandleRegister)
	router.HandleFunc("POST /login", h.HandleLogin)
	router.HandleFunc("DELETE /delete", auth.WithJWTAuth(h.HandleUserDelete, h.store))
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := types.ValidateRegisterUserPayload(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if _, err := h.store.GetUserById(payload.Id); err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %s already exists", payload.Id))
		return
	}

	hashedPass, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		Id:             payload.Id,
		HashedPassword: hashedPass,
		Nickname:       payload.Nickname,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, nil); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := types.ValidateLoginUserPayload(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	currUser, err := h.store.GetUserById(payload.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %s doesn't exist", payload.Id))
		return
	}

	err = auth.ComparePasswords(currUser.HashedPassword, payload.Password)
	if err == auth.IncorrectPassword {
		utils.WriteError(w, http.StatusNotAcceptable, err)
		return
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	tokenString, err := auth.CreateJWT([]byte(config.Envs.JWTSecret), payload.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, map[string]string{"token": tokenString}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) HandleUserDelete(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())

	if err := h.store.DeleteUserById(userId); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, nil); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
