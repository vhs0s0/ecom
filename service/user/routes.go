package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vhs0s0/ecom/types"
	"github.com/vhs0s0/ecom/utils"
)

type Handler struct {
	store *types.UserStore
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get JSON Payload
	var payload types.RegisterUserPayload

	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// Check if the user exists

}
