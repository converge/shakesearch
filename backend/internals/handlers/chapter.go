package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"pulley.com/shakesearch/internals/domain"
	"pulley.com/shakesearch/internals/services"
	"pulley.com/shakesearch/pkg"
	"strconv"
)

type ChapterHandler struct {
	service *services.ChapterService
}

// NewChapterHandler creates a new ChapterHandler, and receives an interface to the ChapterService. The
// interface allow mocking the service during testing, and allow the handler to be loosely coupled to the service.
func NewChapterHandler(service *services.ChapterService) *ChapterHandler {
	return &ChapterHandler{
		service: service,
	}
}

// Search is the handler for the search endpoint. It receives a query parameter, and returns a list of search results.
// A successful response is a JSON array.
func (h *ChapterHandler) Search(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()["query"]

	if len(query[0]) == 0 {
		log.Warn().Msg("unable to search due to missing query parameter in URL")
		http.Error(w, "query parameter not found in URL", http.StatusBadRequest)
		return
	}
	log.Debug().Msgf("searching for query: %s", query[0])

	chapter, err := domain.NewChapter(query[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listChapters, err := h.service.Search(*chapter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", pkg.JSONContentType)
	jsonContent, err := json.Marshal(listChapters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonContent)
}

// GetChapterById is the handler for the get chapter by id endpoint. It receives a chapter id, and returns a chapter in
// JSON format.
func (h *ChapterHandler) GetChapterById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	chapterId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chapter, err := h.service.GetChapterById(chapterId)
	if err == sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonContent, err := json.Marshal(chapter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonContent)
}
