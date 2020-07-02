package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticlest(w http.ResponseWriter, r *http.Request)  {
	articles:=Articles{
		Article{Title: "test title", Desc: "test description", Content: "test content"},
	}
	fmt.Println("Endpoint hit:all articles")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticlest(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Test POST endpoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Home page")
}

func handlerRequests()  {

	myRouter:=mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticlest).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticlest).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main()  {
	handlerRequests()
}
