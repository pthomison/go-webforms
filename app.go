package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
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

func (a *App) init() error {
	db, err := gorm.Open(sqlite.Open("webforms.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	a.db = db

	err = a.db.AutoMigrate(&Message{})
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runServer() error {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/form", a.formHandler)
	http.HandleFunc("/submit-message", submitMessageHandler)

	http.Handle("/static", http.FileServer(http.Dir("./static")))

	fmt.Printf("Starting webserver on %v:%v\n", HOST, PORT)
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		return err
	}

	return nil
}
