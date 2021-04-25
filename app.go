package main

import (
	"fmt"
	utils "github.com/pthomison/golang-utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"strings"
)

type App struct {
	db *gorm.DB
}

type Message struct {
	gorm.Model
	Body  string
	Email string
	Name  string
}

func (a *App) runServer() error {
	var err error

	a.db, err = gorm.Open(sqlite.Open("webforms.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = a.db.AutoMigrate(&Message{})
	if err != nil {
		return err
	}

	http.Handle("/", http.RedirectHandler("/form", 302))
	http.HandleFunc("/form", a.formHandler)
	http.HandleFunc("/submit-message", a.submitMessageHandler)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	fmt.Printf("Starting webserver on %v:%v\n", HOST, PORT)
	err = http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) formHandler(w http.ResponseWriter, r *http.Request) {
	messages := []Message{}

	result := a.db.Find(&messages)
	utils.Check(result.Error)

	t, err := template.ParseFiles("./html/form.html")
	utils.Check(err)

	t.Execute(w, messages)
}

func (a *App) submitMessageHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	utils.Check(err)
	fmt.Printf("post recieved: %+v\n", r)

	result := a.db.Create(&Message{
		Name:  strings.Join(r.Form["name"], ""),
		Email: strings.Join(r.Form["email"], ""),
		Body:  strings.Join(r.Form["message"], ""),
	})
	utils.Check(result.Error)

	http.Redirect(w, r, "/form", 302)
}
