package handler

import (
	"encoding/json"
	"golink/internal/config"
	errapp "golink/internal/errApp"
	"golink/internal/model"
	"golink/internal/service"
	"net/http"
	"uuid"
)

type LinkHandler struct {
	service service.LinkServiceInterface
	env     *config.ENV
}

func NewLinkHandler(service service.LinkServiceInterface, env *config.ENV) *LinkHandler {
	return &LinkHandler{service: service, env: env}
}

func (h *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idString := r.Context().Value(h.env.USER_SUB_KEY).(string)
	id, err := uuid.Parse(idString)
	if err != nil {
		errRes := errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	linkRequest := model.LinkRequest{}

	if err := json.NewDecoder(r.Body).Decode(&linkRequest); err != nil {
		errRes := errapp.NewBadRequestError("Erro no json: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	linkResponse, errRes := h.service.CreateLink(id, linkRequest)
	if errRes != nil {
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(linkResponse)
}

func (h *LinkHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("title")

	linkResponse, errRes := h.service.Redirect(title)
	if errRes != nil {
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	http.Redirect(w, r, linkResponse.URL, http.StatusFound)
}

func (h *LinkHandler) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		errRes := errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	linkUpdate := model.LinkUpdate{}

	if err := json.NewDecoder(r.Body).Decode(&linkUpdate); err != nil {
		errRes := errapp.NewBadRequestError("Erro no json: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	linkResponse, errRes := h.service.Update(id, &linkUpdate)
	if errRes != nil {
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(linkResponse)
}
