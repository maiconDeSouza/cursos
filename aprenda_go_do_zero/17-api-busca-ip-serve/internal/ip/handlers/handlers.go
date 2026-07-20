package handlersIP

import (
	"api-busca-ip-serve/internal/ip/interfaces"
	"api-busca-ip-serve/internal/ip/models"

	"encoding/json"
	"net/http"
)

type HandlersIP struct {
	service interfaces.Services
}

func (h HandlersIP) GetIP(w http.ResponseWriter, r *http.Request) {
	reqIP := models.RequestIP{}

	if err := json.NewDecoder(r.Body).Decode(&reqIP); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := h.service.GetIp(reqIP.HostName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func NewHandlers(service interfaces.Services) *HandlersIP {
	return &HandlersIP{
		service: service,
	}
}
