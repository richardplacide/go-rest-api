package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// HomeHandler handles the Home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there, you're Home")
}

// ARTICLES SECTION -------------------------------------
type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Global array to simulate a DB
var Articles []Article

func populateArticles() {
	Articles = []Article{
		Article{Id: "1", Title: "First article", Desc: "A first article", Content: "Article Content"},
		Article{Id: "2", Title: "Second article", Desc: "This is article 2", Content: "Article Content"},
	}
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}

// END ARTICLES SECTION ------------------------------------

func main() {

	populateArticles()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/articles", returnAllArticles)
	r.HandleFunc("/articles/{id}", returnSingleArticle)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("vue"))))
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting go web server on localhost port 8000")

	log.Fatal(srv.ListenAndServe())
}
