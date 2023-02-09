package handlers

import (
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
	"pulley.com/shakesearch/internals/repositories"
	"pulley.com/shakesearch/internals/services"
	"testing"
)

func TestNewChapterHandler_WithoutService(t *testing.T) {

	chapterService := &services.ChapterService{}
	handler := NewChapterHandler(chapterService)

	if handler == nil {
		t.Errorf("NewChapterHandler() = %v, want %v", handler, chapterService)
	}

}

// todo: merge both tests into one table test
// TestNewChapterHandler_Success is a test for the ChapterHandler, it tests the search endpoint.
// it checks for:
// 1. if the request is successful
// 2. if the request is not successful
func TestNewChapterHandler_Success(t *testing.T) {

	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	dbRepository := repositories.NewPostgresRepository(conn)
	chapterService := &services.ChapterService{
		Store: dbRepository,
	}
	handler := NewChapterHandler(chapterService)

	if handler == nil {
		t.Errorf("NewChapterHandler() = %v, want %v", handler, chapterService)
	}
	r := httptest.NewRequest(http.MethodGet, "/v1/chapter/search?query=hello", nil)
	w := httptest.NewRecorder()

	rows := sqlmock.NewRows([]string{"id", "title", "chapter"})

	mock.ExpectQuery("SELECT").
		WithArgs("hello").WillReturnRows(rows)

	handler.Search(w, r)

	got := w.Result()

	if got.StatusCode != http.StatusOK {
		t.Errorf("NewChapterHandler() = %v, want %v", got.StatusCode, 200)
	}

}

// TestNewChapterHandler_BadRequest is a test for the ChapterHandler, it tests the search endpoint.
// it checks for:
// 1. if the request is not successful because of a bad request when a query string is not provided
func TestNewChapterHandler_BadRequest(t *testing.T) {

	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	dbRepository := repositories.NewPostgresRepository(conn)
	chapterService := &services.ChapterService{
		Store: dbRepository,
	}
	handler := NewChapterHandler(chapterService)

	if handler == nil {
		t.Errorf("NewChapterHandler() = %v, want %v", handler, chapterService)
	}
	r := httptest.NewRequest(http.MethodGet, "/v1/chapter/search?query=", nil)
	w := httptest.NewRecorder()

	rows := sqlmock.NewRows([]string{"id", "title", "chapter"})

	mock.ExpectQuery("SELECT").
		WithArgs("hello").WillReturnRows(rows)

	handler.Search(w, r)

	got := w.Result()

	if got.StatusCode != http.StatusBadRequest {
		t.Errorf("NewChapterHandler() = %v, want %v", got.StatusCode, 200)
	}

}
