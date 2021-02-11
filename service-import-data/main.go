package main

import (
	"log"
	"service-import-data/server"
)



func main() {
	log.Println("Create new httpServer")
	http := server.NewServer()
	http.NewRoutes()
	http.StartAPI()
}
