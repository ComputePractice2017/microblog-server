package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/practice/microblog-server/model"
)

//запускаем сервер
func Run() {
	log.Println("Connecting to rethinkDB on localhost...")
	err := model.InitSesson()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	r := mux.NewRouter()
	r.HandleFunc("/", newTwittHandler).Methods("POST")
	r.HandleFunc("/", getNewsHandler).Methods("GET")
	r.HandleFunc("/", newTwittHandler).Methods("POST")
	r.HandleFunc("/", firstOptionsHandler).Methods("OPTIONS")
	r.HandleFunc("/mexos", getMyTwittsHandler).Methods("GET")
	r.HandleFunc("/mexos/{guid}", deleteTwittHandler).Methods("DELETE")
	r.HandleFunc("/mexos/{guid}", secondOptionsHandler).Methods("OPTIONS")



	log.Println("Running the server on port 8000...")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "DELETE", "OPTIONS"})
  
	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}