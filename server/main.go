package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageTitle struct {
	Title string
	Todo  []*Todo
}

func todo(w http.ResponseWriter, r *http.Request) {

	tododata := PageTitle{Title: "Todo List",
		Todo: []*Todo{
			{Item: "Install Golang", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "Build Go site for web", Done: false},
		},
	}

	tmpl.Execute(w, tododata)
}

func main() {

	fmt.Println("You have launched the server at http://localhost:8080/todo")
	fmt.Println("you should be able to click the adress directly to launch the app in a browser")
	fmt.Println("If you want to end the server just press Ctrl + c in the terminal and it will shut down the server")

	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("./templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./../static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)

	log.Fatal(http.ListenAndServe(":8080", mux))

	fmt.Println("You have launched the server at http://localhost:8080/todo")
	fmt.Println("you should be able to click the adress directly to launch the app in a browser")
}
