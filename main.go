package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.tpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
	}
}

type Data struct {
	VpsID    string
	MgNumber string
}

func outputHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	r.ParseForm()
	t, _ := template.ParseFiles("output.tpl")

	vpsInfo := Data{
		VpsID:    r.FormValue("vpsID"),
		MgNumber: r.FormValue("MGN"),
	}

	t.Execute(w, vpsInfo)

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Post("/output", outputHandler)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", r)
}
