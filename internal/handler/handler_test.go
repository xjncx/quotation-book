package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/xjncx/quotation-book/internal/model"
)

type mockQuoteService struct {
	DeleteFunc func(ctx context.Context, id string) (bool, error)
}

func (m *mockQuoteService) DeleteByID(ctx context.Context, id string) (bool, error) {
	return m.DeleteFunc(ctx, id)
}

func Test_HandleDeleteByID_Success(t *testing.T) {
	qS := &mockQuoteService{
		DeleteFunc: func(ctx context.Context, id string) (bool, error) {
			return true, nil
		},
	}

	h := NewHandler(qS)

	req := httptest.NewRequest(http.MethodDelete, "/quotes/some-id", nil)
	rec := httptest.NewRecorder()
	
	router := mux.NewRouter()
	router.HandleFunc("/quotes/{id}", h.HandleDeleteByID)
	
	router.ServeHTTP(rec, req)
	
	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", rec.Code)
	}
	
	expectedBody := `{"status":"deleted"}` + "\n"
	if rec.Body.String() != expectedBody {
		t.Errorf("unexpected body: got %q, want %q", rec.Body.String(), expectedBody)
	}
	
}

func (m *mockQuoteService) Create(ctx context.Context, q *model.Quote) (*model.Quote, error) {
	panic("not implemented")
}

func (m *mockQuoteService) GetAll(ctx context.Context) ([]*model.Quote, error) {
	panic("not implemented")
}

func (m *mockQuoteService) GetByAuthor(ctx context.Context, author string) ([]*model.Quote, error) {
	panic("not implemented")
}

func (m *mockQuoteService) GetRandom(ctx context.Context) (*model.Quote, error) {
	panic("not implemented")
}
