package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/practice/microblog-server/model"
)

func getMyTwittsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	twitts, err := model.GetMyTwitts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err = json.NewEncoder(w).Encode(twitts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
}

func getNewsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	twitts, err := model.GetNews()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if err = json.NewEncoder(w).Encode(twitts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
}

func firstOptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "OPTIONS, GET, POST")
	w.WriteHeader(http.StatusOK)
}
func secondOptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
}