package pg

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/xjncx/quotation-book/internal/model"
)

type Repository struct {
	conn *sqlx.DB
}

func NewRepository() *Repository {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, pass, dbname, host, port)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{conn: db}
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
		return nil, err
	}

	return &res, nil
}