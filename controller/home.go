package controller

import (
	"html/template"
	"net/http"

	"github.com/prakriti-yan/webshop/viewmodel"
)

type home struct {
	homeTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vw := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vw)
}
