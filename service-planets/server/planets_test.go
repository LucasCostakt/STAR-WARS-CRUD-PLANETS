package server

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestInsertPlanet(t *testing.T) {
	myStructTestResponse := []struct {
		name   string
		id     string
		client http.Client
		want   string
	}{
		{client: http.Client{},
			name: "teste",
			id:   `{"nome":"Yavin IV","clima":"temperate, tropical","terreno":"jungle, rainforests"}`,
			want: `Sucesso ao Inserir um Novo Planeta`,
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodPost, "http://localhost:5000/planets", bytes.NewBuffer([]byte(tt.id)))
			req.Header.Set("Content-Type", "application/json")
			response, _ := tt.client.Do(req)
			got, _ := ioutil.ReadAll(response.Body)

			AssertResponsebody(t, string(got), string(tt.want))
		})
	}
}

func TestGetPlanetById(t *testing.T) {
	myStructTestResponse := []struct {
		name   string
		id     string
		client http.Client
		want   string
	}{
		{client: http.Client{},
			name: "teste",
			id:   `{"id":"60611fcdbe99b18f0b42a843"}`,
			want: `{"id":"60611fcdbe99b18f0b42a843","nome":"alderaan","clima":"temperate","terreno":"grasslands, mountains","films":2}`,
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodGet, "http://localhost:5000/planets", bytes.NewBuffer([]byte(tt.id)))
			req.Header.Set("Content-Type", "application/json")
			response, _ := tt.client.Do(req)
			got, _ := ioutil.ReadAll(response.Body)

			AssertResponsebody(t, string(got), string(tt.want))
		})
	}
}

func TestGetPlanetByName(t *testing.T) {
	myStructTestResponse := []struct {
		name     string
		nameSend string
		client   http.Client
		want     string
	}{
		{client: http.Client{},
			name:     "teste",
			nameSend: `{"nome":"Alderaan"}`,
			want:     `{"id":"60611fcdbe99b18f0b42a843","nome":"alderaan","clima":"temperate","terreno":"grasslands, mountains","films":2}`,
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodGet, "http://localhost:5000/planets", bytes.NewBuffer([]byte(tt.nameSend)))
			req.Header.Set("Content-Type", "application/json")
			response, _ := tt.client.Do(req)
			got, _ := ioutil.ReadAll(response.Body)

			AssertResponsebody(t, string(got), string(tt.want))
		})
	}
}

func TestGetAllPlanet(t *testing.T) {
	myStructTestResponse := []struct {
		name   string
		client http.Client
		want   string
	}{
		{client: http.Client{},
			name: "teste",
			want: `[{"id":"60611fcdbe99b18f0b42a843","nome":"alderaan","clima":"temperate","terreno":"grasslands, mountains","films":2}]`,
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodGet, "http://localhost:5000/planets", nil)
			req.Header.Set("Content-Type", "application/json")
			response, _ := tt.client.Do(req)
			got, _ := ioutil.ReadAll(response.Body)

			AssertResponsebody(t, string(got), string(tt.want))
		})
	}

}

func TestDeletePlanet(t *testing.T) {
	myStructTestResponse := []struct {
		name   string
		id     string
		client http.Client
		want   string
	}{
		{client: http.Client{},
			name: "teste",
			id:   `{"id":"6061519aece1bbff2d15632f"}`,
			want: `Documento deletado com Sucesso!`,
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodDelete, "http://localhost:5000/planets", bytes.NewBuffer([]byte(tt.id)))
			req.Header.Set("Content-Type", "application/json")
			response, _ := tt.client.Do(req)
			got, _ := ioutil.ReadAll(response.Body)

			AssertResponsebody(t, string(got), string(tt.want))
		})
	}

}

func AssertResponsebody(t *testing.T, got, expectedResponse string) {
	t.Helper()
	if got != expectedResponse {
		t.Errorf("body is wrong, got %q want %q\n", got, expectedResponse)
	}
}
