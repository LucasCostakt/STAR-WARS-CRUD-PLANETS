package server

type Storage interface {
	NewRoutes()
	StartAPI()
}