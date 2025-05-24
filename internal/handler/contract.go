package handler

import (
	"context"

	"github.com/xjncx/quotation-book/internal/model"
)

type QuoteService interface {
	Create(ctx context.Context, quote *model.Quote) (*model.Quote, error)
	GetAll(ctx context.Context) ([]*model.Quote, error)
	GetByAuthor(ctx context.Context, author string) ([]*model.Quote, error)
	GetRandom(ctx context.Context) (*model.Quote, error)
	DeleteByID(ctx context.Context, id string) (bool, error)
}
