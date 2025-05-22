package service

import (
	"context"

	"github.com/xjncx/quotation-book/internal/model"
)

type Service struct {
	repo Repo
}

func New(r Repo) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, quote *model.Quote) (*model.Quote, error) {
	created, err := s.repo.Insert(ctx, quote)
	if err != nil {
		return quote, err
	}
	return created, nil
}
