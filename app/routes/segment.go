package routes

import (
	"challange/app/controller"
	"net/http"
)

type SegmentRoutes struct {
	segmentController controller.SegmentController
}

func NewSegmentRoutes(sc controller.SegmentController) SegmentRoutes {
	return SegmentRoutes{
		segmentController: sc,
	}
}

func (r SegmentRoutes) AddRoutes(sm *http.ServeMux) {
	sm.HandleFunc("/users", r.segmentController.ListCreate)
}
