package handler

type CreateQuoteRequest struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type QuoteResponse struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}