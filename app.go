package main

import (
	"fmt"
	utils "github.com/pthomison/golang-utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"strconv"
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

type templateData struct {
	Messages []Message
	Stats    struct {
		MessageCount int
	}
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

	for i := 0; i < 100; i++ {
		result := a.db.Create(&Message{
			Name:  "test user",
			Email: "mail@example.com",
			Body:  "message number: " + strconv.Itoa(i),
		})
		utils.Check(result.Error)
	}

	http.Handle("/", http.RedirectHandler("/form", 302))
	http.HandleFunc("/form", a.formHandler)
	http.HandleFunc("/submit-message", a.submitMessageHandler)
	// strip prefix is required for MIME type
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	fmt.Printf("Starting webserver on %v\n", ADDR)
	err = http.ListenAndServe(ADDR, nil)
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

	t.Execute(w, &templateData{
		Messages: messages,
	})
}

func (a *App) submitMessageHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	utils.Check(err)

	result := a.db.Create(&Message{
		Name:  strings.Join(r.Form["name"], ""),
		Email: strings.Join(r.Form["email"], ""),
		Body:  strings.Join(r.Form["message"], ""),
	})
	utils.Check(result.Error)

	http.Redirect(w, r, "/form", 302)
}
