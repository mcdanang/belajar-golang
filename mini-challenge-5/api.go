package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Address  string
	Job      string
	Reason   string
}

var users = []User{
	{ID: 1, Name: "John Doe", Email: "john.doe@example.com", Password: "asd123", Address: "123 Main St, Cityville, USA", Job: "Software Engineer", Reason: "Sunshine and warmth"},
	{ID: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Password: "asd123", Address: "234 Main St, Cityville, USA", Job: "Backend Engineer", Reason: "Money, money, and money"},
	{ID: 3, Name: "Michael Johnson", Email: "michael.johnson@example.com", Password: "asd123", Address: "345 Main St, Cityville, USA", Job: "Frontend Engineer", Reason: "K-Pop Idol"},
	{ID: 4, Name: "Sarah Wilson", Email: "sarah.wilson@example.com", Password: "asd123", Address: "456 Main St, Cityville, USA", Job: "Devops Engineer", Reason: "Manga & Anime"},
	{ID: 5, Name: "Robert Jackson", Email: "robert.jackson@example.com", Password: "asd123", Address: "567 Main St, Cityville, USA", Job: "UI/UX Designer", Reason: "Gadgets"},
}

var PORT = ":9090"

func main() {
	http.HandleFunc("/", getUsers)
	// http.HandleFunc("/users", getUsers)
	http.HandleFunc("/user", getUser)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, users)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user := User{}
		for _, v := range users {
			if v.Email == email {
				if v.Password == password {
					user = v
					break
				}
				sendErrorMessage("Wrong Password", w)
				return
			}
			sendErrorMessage("Email is not registered!", w)
			return
		}

		tpl, err := template.ParseFiles("user.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, user)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func sendErrorMessage(message string, w http.ResponseWriter) {
	tpl, err := template.ParseFiles("error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, message)
	return
}
