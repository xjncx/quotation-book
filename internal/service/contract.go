package service

import (
	"context"

	"github.com/xjncx/quotation-book/internal/model"
)

type Repo interface {
	Insert(ctx context.Context, quote *model.Quote) (*model.Quote, error)
	GetAll(ctx context.Context) ([]*model.Quote, error)
	GetRandom(ctx context.Context) (*model.Quote, error)
	FindByAuthor(ctx context.Context, author string) ([]*model.Quote, error)
	DeleteByID(ctx context.Context, id string) (bool, error)
}
