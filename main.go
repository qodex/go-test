package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var articles []Article

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	var article Article
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &article)
	articles = append(articles, article)
	json, _ := json.Marshal(articles)
	fmt.Fprintf(w, string(json))
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	articleID := mux.Vars(r)["id"]
	var article Article
	for _, a := range articles {
		if a.ID == articleID {
			article = a
		}
	}
	json, _ := json.Marshal(article)
	fmt.Fprintf(w, string(json))
}

func findByTagAndDate(w http.ResponseWriter, r *http.Request) {
	tagParam := mux.Vars(r)["tag"]
	dateParam := mux.Vars(r)["date"]

	response := FindResponse{}
	response.Tag = tagParam
	response.Count = 0
	response.Articles = make([]string, 0, 100)
	response.RelatedTags = make([]string, 0, 100)

	for _, a := range articles {
		dateStr := strings.ReplaceAll(a.Date, "-", "")
		if dateStr == dateParam {
			response.RelatedTags = MergeUnique(response.RelatedTags, a.Tags, tagParam)
			if Contains(a.Tags, tagParam) {
				response.Count = response.Count + 1
				response.Articles = append(response.Articles, a.ID)
			}
		}
	}

	json, _ := json.Marshal(response)

	fmt.Fprintf(w, string(json))
}

func main() {
	articles = make([]Article, 0, 1000)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/articles", articlesHandler).Methods("POST")
	router.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/tags/{tag}/{date}", findByTagAndDate).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
