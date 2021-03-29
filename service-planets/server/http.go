package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"service-consult/repository"
	"strings"
)

type httpServer struct {
	http.Handler
}

func NewServer() Storage {
	return new(httpServer)
}

func (h *httpServer) NewRoutes() {
	log.Println("Iniciando Rotas")

	client, _ := repository.NewMongoConnect()
	serv := NewService(client)

	router := http.NewServeMux()
	router.Handle("/planets", http.HandlerFunc(serv.PlanetsOperator))

	h.Handler = router
}

func (h *httpServer) StartAPI() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Println("API Iniciada")
	log.Println("** Serviço inciado na Porta " + port + " **")
	if err := http.ListenAndServe(":"+port, h); err != nil {
		log.Fatal("init server error in StartApi(), ", err)
	}
}

func (serv *Server) PlanetsOperator(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		var response []byte
		filters := &repository.Filters{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Read request body error in PlanetsOperator(): ", err)
		}
		json.Unmarshal(body, filters)

		if filters.Name != "" {
			response, err = serv.Repository.CounsultPlanetByName(strings.ToLower(filters.Name))
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

		} else if filters.Id != "" {
			response, err = serv.Repository.CounsultPlanetByID(filters.Id)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

		} else {
			response, err = serv.Repository.CounsultAllPlanets()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

		}
		w.Write(response)

	case http.MethodPost:
		newInsertPlanet := repository.NewInsertPlanet{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Erro ao ler request PlanetsOperator: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.Unmarshal(body, &newInsertPlanet)
		countPlanet, err := PlanetCount(newInsertPlanet.Nome)
		if err != nil {
			log.Println("Erro ao ler request PlanetsOperator: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newInsertPlanet.Nome = strings.ToLower(newInsertPlanet.Nome)
		newInsertPlanet.FilmsCount = countPlanet
		response, err := serv.Repository.InsertNewPlanet(newInsertPlanet)
		w.Write(response)

	case http.MethodDelete:
		filters := &repository.Filters{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Erro ao ler request PlanetsOperator: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.Unmarshal(body, filters)
		err = serv.Repository.DeletePlanetById(filters.Id)
		if err != nil {
			log.Println("Erro DeletePlanetById() PlanetsOperator: ", err)
			w.WriteHeader(http.StatusNotModified)
			return
		}
		w.Write([]byte("Documento deletado com Sucesso!"))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
