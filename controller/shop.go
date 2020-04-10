package controller

import (
	"html/template"
	"net/http"

	"github.com/prakriti-yan/webshop/viewmodel"
)

type shop struct {
	shopTemplate *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	vw := viewmodel.NewShop()
	s.shopTemplate.Execute(w, vw)
}
