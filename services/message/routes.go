package message

import (
	"net/http"
	"strconv"

	"github.com/kasariks/simple_messenger_api/services/auth"
	"github.com/kasariks/simple_messenger_api/types"
	"github.com/kasariks/simple_messenger_api/utils"
)

type Handler struct {
	store     types.MessageStore
	userStore types.UserStore
}

func NewHandler(store types.MessageStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:     store,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /getmessages", auth.WithJWTAuth(h.HandleGetAllMessages, h.userStore))
	router.HandleFunc("POST /sendmessage", auth.WithJWTAuth(h.HandleSendMessage, h.userStore))
	router.HandleFunc("GET /getgottenmessages", auth.WithJWTAuth(h.HandleGetGottenMessages, h.userStore))
}

func (h *Handler) HandleGetAllMessages(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	messages, err := h.store.GetMessagesByAuthorId(userId, page)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, messages); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) HandleSendMessage(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())

	var payload types.SendMessagePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := types.ValidateSendMessagePayload(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateMessage(types.Message{
		AuthorId:    userId,
		RecipientId: payload.RecipientId,
		Value:       payload.Value,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) HandleGetGottenMessages(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	messages, err := h.store.GetMessagesByRecipientId(userId, page)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, messages); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
