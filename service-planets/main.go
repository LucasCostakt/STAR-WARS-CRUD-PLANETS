package main

import (
	"log"
	"service-consult/server"
)

func main() {
	log.Println("Criado Novo httpServer")
	http := server.NewServer()
	http.NewRoutes()
	http.StartAPI()
}
