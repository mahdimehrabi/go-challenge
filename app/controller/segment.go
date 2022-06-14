package controller

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"challange/app/repository"
	"fmt"
	"net/http"
)

type SegmentController struct {
	logger         interfaces.Logger
	userRepository repository.UserRepository
}

func NewSegmentController(
	logger infrastracture.SegmentLogger,
	userRepository repository.UserRepository) SegmentController {
	return SegmentController{
		logger:         &logger,
		userRepository: userRepository,
	}
}

func (c SegmentController) ListCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Print("Dgsgds")
	}
}
