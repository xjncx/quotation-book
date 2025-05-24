package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
	"github.com/xjncx/quotation-book/internal/model"
	"github.com/xjncx/quotation-book/internal/repository"
)

type Repository struct {
	conn *sql.DB
}

func NewRepository(dsn string) (*Repository, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("NewRepository: opened connection failed %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("NewRepository: ping failed: %w", err)
	}

	return &Repository{conn: db}, nil
}

func (r *Repository) Insert(ctx context.Context, quote *model.Quote) (*model.Quote, error) {
	query := `
	INSERT INTO quotes (quote_uuid, author, quote_text) 
	VALUES (uuid_generate_v4(), $1, $2) 
	RETURNING id, quote_uuid, author, quote_text;`

	row := r.conn.QueryRowContext(ctx, query, quote.Author, quote.Text)

	var res model.Quote
	err := row.Scan(&res.ID, &res.UUID, &res.Author, &res.Text)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			log.Printf("Insert: duplicate quote: %v", pqErr.Detail)
			return nil, repository.ErrDuplicate
		}

		log.Printf("Insert: query failed: %v", err)
		return nil, fmt.Errorf("Insert: failed to scan row: %w", err)
	}

	return &res, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]*model.Quote, error) {
	query := `
	SELECT id, quote_uuid, author, quote_text
	FROM quotes
	ORDER BY created_at DESC;`

	var quotes []*model.Quote

	rows, err := r.conn.QueryContext(ctx, query)

	if err != nil {
		log.Printf("GetAll: query failed: %v", err)
		return nil, fmt.Errorf("GetAll: failed to fetch quotes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var q model.Quote
		if err := rows.Scan(&q.ID, &q.UUID, &q.Author, &q.Text); err != nil {
			log.Printf("GetAll: scan failed: %v", err)
			return nil, fmt.Errorf("GetAll: failed to scan row: %w", err)
		}
		quotes = append(quotes, &q)
	}

	return quotes, nil
}

func (r *Repository) FindByAuthor(ctx context.Context, author string) ([]*model.Quote, error) {
	query := `
	SELECT id, quote_uuid, author, quote_text
	FROM quotes
	WHERE author = $1
	ORDER BY created_at DESC;
	`
	rows, err := r.conn.QueryContext(ctx, query, author)

	if err != nil {
		log.Printf("FindByAuthor: query failed: %v", err)
		return nil, fmt.Errorf("FindByAuthor: failed to fetch quotes: %w", err)
	}
	defer rows.Close()

	var quotes []*model.Quote

	for rows.Next() {
		var q model.Quote
		if err := rows.Scan(&q.ID, &q.UUID, &q.Author, &q.Text); err != nil {
			log.Printf("FindByAuthor: scan failed: %v", err)
			return nil, fmt.Errorf("FindByAuthor: failed to scan row: %w", err)
		}
		quotes = append(quotes, &q)
	}

	if len(quotes) == 0 {
		log.Printf("FindByAuthor: no quotes found for author=%q", author)
		return nil, repository.ErrNotFound
	}

	return quotes, nil
}

func (r *Repository) GetRandom(ctx context.Context) (*model.Quote, error) {
	query := `
	SELECT id, quote_uuid, author, quote_text
	FROM quotes
	ORDER BY RANDOM()
	LIMIT 1;
	`
	row := r.conn.QueryRowContext(ctx, query)

	var res model.Quote
	err := row.Scan(&res.ID, &res.UUID, &res.Author, &res.Text)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("GetRandom: no quotes found")
			return nil, repository.ErrNotFound
		}
		log.Printf("GetRandom: query failed: %v", err)
		return nil, fmt.Errorf("GetRandom: failed to scan row: %w", err)
	}

	return &res, nil
}

func (r *Repository) DeleteByID(ctx context.Context, id string) (bool, error) {
	query := `DELETE FROM quotes WHERE quote_uuid = $1;`

	res, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("DeleteByID: query failed: %v", err)
		return false, fmt.Errorf("DeleteByID: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("DeleteByID: rows affected failed: %v", err)
		return false, fmt.Errorf("DeleteByID: %w", err)
	}

	if count == 0 {
		log.Printf("DeleteByID: no quote found with id=%s", id)
		return false, repository.ErrNotFound
	}

	return true, nil
}
