package handler

import (
	"pasarwarga/Users"
)

type SessionHandler struct {
	UserService Users.Service
}

func NewSessionHandler(UserService Users.Service) *SessionHandler {
	return &SessionHandler{UserService}
}
