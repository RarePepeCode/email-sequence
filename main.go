package main

import (
	"github.com/RarePepeCode/email-sequence/middleware"
)

func main() {
	srv := middleware.StartServer()
	srv.ListenAndServe()
}
