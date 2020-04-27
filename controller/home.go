package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/prakriti-yan/webshop/model"
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
	// server push:
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}
	vw := viewmodel.NewHome()
	w.Header().Add("Content-Type", "text/html")
	// time.Sleep(4 * time.Second)
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
		if user, err := model.Login(email, password); err == nil {
			log.Printf("User has logged in: %v\n", user)
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			vw.Email = email
			vw.Password = password
		}
	}
	w.Header().Add("Content-Type", "text/html")
	// when we start adding gzip compression, the data is gonna come back compressed and we
	// need to give the content type hint in order for the browser to understand what type of data
	// it is getting!
	h.loginTemplate.Execute(w, vw)
}
