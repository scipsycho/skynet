package server

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("frontPage.html", "login.html", "signup.html", "afterLogin.html", "createRecord.html", "accessAttributes.html"))

func renderTemplate(w http.ResponseWriter, name string, val interface{}) {
	err := templates.ExecuteTemplate(w, name+".html", val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "createRecord", nil)
}

func frontPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "frontPage", nil)
}

func displaySignUpHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "signup", nil)
}

func displayLoginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login", nil)
}

func randomDisplay(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "afterLogin", nil)
}
