package controller

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"challange/app/services"
	"io/ioutil"
	"net/http"
)

type SegmentController struct {
	logger         interfaces.Logger
	segmentService services.SegmentService
}

func NewSegmentController(
	logger infrastracture.SegmentLogger,
	segmentService services.SegmentService) SegmentController {
	return SegmentController{
		logger:         &logger,
		segmentService: segmentService,
	}
}

func (c SegmentController) ListCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			infrastracture.BadRequestResponse(w)
			return
		}
		if err = c.segmentService.CreateUser(b); err != nil {
			infrastracture.ErrorResponse(err, c.logger, w)
			return
		}

		infrastracture.SuccessResponse(w, "user created successfully")
	}
}
