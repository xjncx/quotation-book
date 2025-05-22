package handler

import (
	"context"

	"github.com/xjncx/quotation-book/internal/model"
)

type QuoteService interface{
	Create(ctx context.Context, data *model.Quote)(*model.Quote, error)
	GetAll()
	GetRandom()
	GetByAuthor()
	DeleteByID()
}