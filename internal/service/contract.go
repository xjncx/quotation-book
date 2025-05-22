package service

import (
	"context"

	"github.com/xjncx/quotation-book/internal/model"
)

type Repo interface {
	Insert(ctx context.Context, quote *model.Quote) (*model.Quote, error)
}
