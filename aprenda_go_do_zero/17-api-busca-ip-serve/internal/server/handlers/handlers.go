package handlersServer

import (
	"api-busca-ip-serve/internal/server/interfaces"
	"api-busca-ip-serve/internal/server/models"
	"encoding/json"
	"net/http"
)

type HandlersServer struct {
	service interfaces.Services
}

func (h HandlersServer) GetServer(w http.ResponseWriter, r *http.Request) {
	reqServer := models.RequestServer{}

	if err := json.NewDecoder(r.Body).Decode(&reqServer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := h.service.GetServer(reqServer.HostName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func NewHandlers(service interfaces.Services) *HandlersServer {
	return &HandlersServer{
		service: service,
	}
}
