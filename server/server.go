package server

import (
	"log"
	"net/http"
	"os"
)

func Init() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	router := Router()
	log.Printf("Server started @ port %s\n", port)
	err := http.ListenAndServe(":"+port, router)
	log.Fatal(err)
}
