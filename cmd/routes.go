package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// all route handlers must take these parameters in
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home!")
}

// Todo is a todo with a title and content
type Todo struct {
	Title   string
	Content string
}

// PageVariables are variables sent to the html template
type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		PageTitle: "Get todos",
		PageTodos: todos,
	}
	t, err := template.ParseFiles("todos.html") // returns whatever is parsed and an error if there is one

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // throw an error with 400 error code
		log.Print("Template parsing error: ", err)
	}

	err = t.Execute(w, pageVariables) // if no errors, execute what has been parsed
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // grabs everything out of the request

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Request parsing error: ", err)

	}

	todo := Todo{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	todos = append(todos, todo)

	log.Print(todos)
	http.Redirect(w, r, "/todos/", http.StatusSeeOther)
}

func main() {
	// Go uses net/http to handle routes
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo/", addTodo)

	fmt.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // everything inside gets executed regardless
}

// use swapi.dev instead
