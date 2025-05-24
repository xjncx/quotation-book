package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/xjncx/quotation-book/internal/model"
	"github.com/xjncx/quotation-book/internal/repository"
)

type Handler struct {
	qS QuoteService
}

func NewHandler(quoteS QuoteService) *Handler {
	return &Handler{qS: quoteS}
}

func (h *Handler) HandleCreateQuote(w http.ResponseWriter, r *http.Request) {
	var req CreateQuoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if strings.TrimSpace(req.Author) == "" || strings.TrimSpace(req.Quote) == "" {
		writeError(w, http.StatusBadRequest, fmt.Errorf("author and quote must not be empty"))
		return
	}

	quote := &model.Quote{
		Author: req.Author,
		Text:   req.Quote,
	}

	saved, err := h.qS.Create(r.Context(), quote)
	if err != nil {

		if errors.Is(err, repository.ErrDuplicate) {
			writeError(w, http.StatusConflict, errors.New("duplicate quote"))
			return
		}

		writeError(w, http.StatusInternalServerError, err)
		return
	}

	resp := QuoteResponse{
		ID:     saved.UUID,
		Author: saved.Author,
		Quote:  saved.Text,
	}

	writeJSON(w, http.StatusCreated, resp)
}

func (h *Handler) HandleGetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")

	var (
		quotes []*model.Quote
		err    error
	)

	if author != "" {
		quotes, err = h.qS.GetByAuthor(r.Context(), author)

	} else {
		quotes, err = h.qS.GetAll(r.Context())
	}

	if err != nil {
		if errors.Is(err, repository.ErrNotFound) && author != "" {
			writeError(w, http.StatusNotFound, fmt.Errorf("no quotes found for author %q", author))
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	res := make([]QuoteResponse, 0, len(quotes))

	for _, q := range quotes {
		res = append(res, ToQuoteResponse(q))
	}

	log.Printf("Fetched %d quotes (filter: %q)", len(quotes), author)

	writeJSON(w, http.StatusOK, res)
}

func (h *Handler) HandleGetRandomQuote(w http.ResponseWriter, r *http.Request) {

	q, err := h.qS.GetRandom(r.Context())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			writeError(w, http.StatusNotFound, errors.New("no quotes available"))
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	res := ToQuoteResponse(q)
	writeJSON(w, http.StatusOK, res)
}

func (h *Handler) HandleDeleteByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uuid := vars["id"]

	if uuid == "" {
		writeError(w, http.StatusBadRequest, fmt.Errorf("missing 'id' in path"))
		return
	}

	log.Printf("Attempting to delete quote id: %s", uuid)
	_, err := h.qS.DeleteByID(r.Context(), uuid)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			writeError(w, http.StatusNotFound, fmt.Errorf("quote not found"))
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, StatusResponse{Status: "deleted"})
}

func ToQuoteResponse(q *model.Quote) QuoteResponse {
	return QuoteResponse{
		ID:     q.UUID,
		Author: q.Author,
		Quote:  q.Text,
	}
}
