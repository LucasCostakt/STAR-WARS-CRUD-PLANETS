package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"service-consult/repository"
	"strings"
)

type ArrayPlanet struct {
	PP []PlanetFilmsCount `json:"results"`
}

type PlanetFilmsCount struct {
	FilmURLs []string `json:"films"`
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewRequest(method string, url string, requestBody []byte) (*http.Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	return request, err
}

func PlanetCount(planetName string) (int64, error) {

	var countFilms int64

	client := &http.Client{}
	planetName = strings.ToLower(planetName)
	url := repository.UrlPlanets + planetName
	request, err := NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("Erro no Request PlanetCount()")
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("Erro no Response PlanetCount()")
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("Erro na leitura do Body PlanetCount()")
	}

	pl := &ArrayPlanet{}
	json.Unmarshal(body, pl)

	for _, b := range pl.PP {
		countFilms = int64(len(b.FilmURLs))
	}
	return countFilms, nil
}
