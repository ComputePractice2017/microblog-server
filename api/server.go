package api

import (
	"log"
	
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

	//http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}