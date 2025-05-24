package service

import (
	"context"
	"fmt"
	"log"

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
		log.Printf("Create: insert failed: %v", err)
		return nil, fmt.Errorf("Create: failed to repo.Insert: %w", err)

	}
	return created, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*model.Quote, error) {
	fetched, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("GetAll: fetch failed: %v", err)
		return nil, fmt.Errorf("GetAll: failed to repo.GetAll: %w", err)
	}
	return fetched, nil
}

func (s *Service) GetRandom(ctx context.Context) (*model.Quote, error) {
	fetched, err := s.repo.GetRandom(ctx)
	if err != nil {
		log.Printf("GetRandom: fetch failed: %v", err)
		return nil, fmt.Errorf("GetRandom: failed to repo.GetRandom: %w", err)
	}
	return fetched, nil
}

func (s *Service) GetByAuthor(ctx context.Context, author string) ([]*model.Quote, error) {
	fetched, err := s.repo.FindByAuthor(ctx, author)
	if err != nil {
		log.Printf("GetByAuthor: fetch failed: %v", err)
		return nil, fmt.Errorf("GetByAuthor: failed to repo.GetByAuthor: %w", err)
	}
	return fetched, nil
}

func (s *Service) DeleteByID(ctx context.Context, uuid string) (bool, error) {
	res, err := s.repo.DeleteByID(ctx, uuid)
	if err != nil {
		log.Printf("DeleteByID: fetch failed: %v", err)
		return true, fmt.Errorf("DeleteByID: failed to repo.DeleteByID: %w", err)
	}
	return res, nil
}
