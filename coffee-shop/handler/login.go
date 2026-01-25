package handler

import "net/http"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password`
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {

}
