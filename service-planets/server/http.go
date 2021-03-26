package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type httpServer struct {
	http.Handler
}

func NewServer() Storage {
	return new(httpServer)
}

func (h *httpServer) NewRoutes() {
	log.Println("Init Routes")
	router := http.NewServeMux()
	router.Handle("/planets", http.HandlerFunc(planetsOperator))

	h.Handler = router
}

func (h *httpServer) StartAPI() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Println("Start API")
	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, h); err != nil {
		log.Fatal("init server error in StartApi(), ", err)
	}
}

func planetsOperator(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Metodo get")
	case http.MethodPost:
		fmt.Fprintf(w, "Metodo post")
	case http.MethodPut:
		fmt.Fprintf(w, "Metodo put")
	case http.MethodDelete:
		fmt.Fprintf(w, "Metodo delete")
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
