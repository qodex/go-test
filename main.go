package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qodex/articledao"
)

var articlesDao articledao.ArticleDAO

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	var article articledao.Article
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &article)
	articleID, e := articlesDao.SaveArticle(article)
	if e == nil {
		fmt.Fprintf(w, articleID)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, e.Error())
	}
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	articleID := mux.Vars(r)["id"]
	article, e := articlesDao.GetArticle(articleID)
	if e == nil {
		json, _ := json.Marshal(article)
		fmt.Fprintf(w, string(json))
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, e.Error())
	}
}

func findByTagAndDate(w http.ResponseWriter, r *http.Request) {
	tagParam := mux.Vars(r)["tag"]
	dateParam := mux.Vars(r)["date"]
	response, e := articlesDao.FindByTagAndDate(tagParam, dateParam)
	if e == nil {
		json, _ := json.Marshal(response)
		fmt.Fprintf(w, string(json))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, e.Error())
	}
}

func init() {
	fmt.Println("Starting articles service...")
	articlesDao = new(articledao.ArticleDAOInMem)
}

func main() {
	port := ":8080"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/articles", articlesHandler).Methods("POST")
	router.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/tags/{tag}/{date}", findByTagAndDate).Methods("GET")
	fmt.Println("Started on port ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
