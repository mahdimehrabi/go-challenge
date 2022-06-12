package routes

import (
	"fmt"
	"net/http"
)

type SegmentRoutes struct {
	//todo:defining controllers
}

func NewSegmentRoutes() SegmentRoutes {
	return SegmentRoutes{}
}

func (r SegmentRoutes) AddRoutes(sm *http.ServeMux) {
	sm.HandleFunc("/segments", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello segment")
	})
}
