package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func GreetingPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/start.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		message := fmt.Sprintf("Hello, %s", name)

		tmpl, err := template.ParseFiles("static/greeting.html")
		if err != nil {
			http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, map[string]interface{}{"Message": message})
		return
	}

	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", GreetingPage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
