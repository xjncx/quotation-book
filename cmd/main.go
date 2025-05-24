package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/xjncx/quotation-book/internal/handler"
	"github.com/xjncx/quotation-book/internal/repository/pg"
	"github.com/xjncx/quotation-book/internal/service"
)

func main() {

	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, pass, dbname, host, port)

	repo, err := pg.NewRepository(dsn)
	if err != nil {
		log.Fatalf("main: failed to connect to DB: %v", err)
	}

	svc := service.New(repo)
	h := handler.NewHandler(svc)

	r := mux.NewRouter()
	r.Use(handler.Middleware)
	r.HandleFunc("/quotes", h.HandleCreateQuote).Methods("POST")
	r.HandleFunc("/quotes", h.HandleGetQuotes).Methods("GET")
	r.HandleFunc("/quotes/random", h.HandleGetRandomQuote).Methods("GET")
	r.HandleFunc("/quotes/{id}", h.HandleDeleteByID).Methods("DELETE")

	addr := ":8080"
	log.Printf("Server is running on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("main: server failed: %v", err)
	}
}
