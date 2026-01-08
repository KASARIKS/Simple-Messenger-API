package user

import (
	"fmt"
	"net/http"

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
	router.HandleFunc("POST /register", h.HandleLogin)
	router.HandleFunc("DELETE /register", h.HandleUserDelete)
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if _, err := h.store.GetUserById(payload.Id); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %s already exists", payload.Id))
		return
	}

	err := h.store.CreateUser(types.User{
		Id:       payload.Id,
		Password: payload.Password,
		Nickname: payload.Nickname,
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

}

func (h *Handler) HandleUserDelete(w http.ResponseWriter, r *http.Request) {

}
