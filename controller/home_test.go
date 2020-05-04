package controller

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginExecutesCorrectTemplate(t *testing.T) {
	h := new(home)
	// create a new controller!
	expected := "Login page Template"
	h.loginTemplate, _ = template.New("").Parse(expected)
	r := httptest.NewRequest(http.MethodGet, "/login", nil)
	w := httptest.NewRecorder()

	h.handleLogin(w, r)

	actual, _ := ioutil.ReadAll(w.Result().Body)
	// Response of the body can be read using any method that could read data
	// from incoming byte stream. Simplest of them is ReadAll function provided in ioutil package.
	if string(actual) != expected {
		t.Errorf("Failed execute correct template!")
	}
}
