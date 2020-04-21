package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/prakriti-yan/webshop/viewmodel"
)

type home struct {
	homeTemplate  *template.Template
	loginTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vw := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vw)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vw := viewmodel.NewLogin()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error logging in : %v", err))
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		if email == "test@gmail.com" && password == "1234" {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			vw.Email = email
			vw.Password = password
		}
	}
	h.loginTemplate.Execute(w, vw)
}
