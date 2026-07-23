package api

import (
	"net/http"
	"github.com/jacksonfishburn/go-cab/internal/files"
)

type API struct {
	manager files.Manager 
}

func newAPI() API {
	return API{manager: files.Manager{}}
}

func (api API) Ping(w http.ResponseWriter, r *http.Request) {
	if api.manager.Ping() {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (api API) Add(w http.ResponseWriter, r *http.Request) {
	response := api.manager.Add()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (api API) Grab(w http.ResponseWriter, r *http.Request) {
	response := api.manager.Grab()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (api API) Del(w http.ResponseWriter, r *http.Request) {
	response := api.manager.Del()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (api API) Peek(w http.ResponseWriter, r *http.Request) {
	response := api.manager.Peek()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}