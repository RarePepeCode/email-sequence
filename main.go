package main

import (
	"log"

	"github.com/RarePepeCode/email-sequence/middleware"
)

func main() {
	log.Println("init")

	srv := middleware.StartServer()
	srv.ListenAndServe()

	log.Println("end")
}
