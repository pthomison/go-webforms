package main

import (
	"fmt"
	"html/template"
	// "io/ioutil"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		defaultHandler(w, r)
		return
	}

	http.Redirect(w, r, "/form", 302)
	return
}

func (a *App) formHandler(w http.ResponseWriter, r *http.Request) {
	// app := &App{
	// 	Messages: []Message{
	// 		Message{
	// 			Name:  "Patrick Thomison",
	// 			Email: "p.thomison@gmail.com",
	// 			Body:  "This is a test message",
	// 		},
	// 	},
	// }

	t, _ := template.ParseFiles("./html/form.html")
	t.Execute(w, nil)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./static/form.html")
	t.Execute(w, nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./html/default.html")
	t.Execute(w, nil)
}

func submitMessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("post recieved: %+v\n", r)
	http.Redirect(w, r, "/form", 302)
	return
}
