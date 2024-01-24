package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

// Please use Ctrl + Q to see documentation on anything

func main() {
	fmt.Println("Server on")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		fmt.Println(title)
		fmt.Println(director)
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		tmpl, _ := template.New("T").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
