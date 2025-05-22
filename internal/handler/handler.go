package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/xjncx/quotation-book/internal/model"
)

//"github.com/xjncx/quotation-book/model"

type Handler struct {
	qS QuoteService
}

func NewHandler(quoteS QuoteService) *Handler {
	return &Handler{qS: quoteS}
}

func (h *Handler) HandleRequest(
	w http.ResponseWriter,
	r *http.Request,
	data any,
	action func(any) (any, error)) {

	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	log.Printf("Received data: %+v\n", data)

	//

	resp, err := action(data)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	writeJSON(w, http.StatusCreated, resp)
}

func (h *Handler) HandleCreateQuote(w http.ResponseWriter, r *http.Request) {
	var q model.Quote

	h.HandleRequest(w, r, &q, func(input any) (any, error) {
		return h.qS.Create(r.Context(), input.(*model.Quote))
	})

}
