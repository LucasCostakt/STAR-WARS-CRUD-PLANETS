package server

import (
	"net/http"
	"service-consult/repository"
)

type Storage interface {
	NewRoutes()
	StartAPI()
}

type Server struct {
	Repository repository.Repository
}

type ServerList interface {
	PlanetsOperator(w http.ResponseWriter, r *http.Request)
}

//NewDocumentService creates a new instance of DocumentService
func NewService(repo repository.Repository) ServerList {
	return &Server{
		Repository: repo,
	}
}
