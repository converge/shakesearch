package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"pulley.com/shakesearch/internals/domain"
	"testing"
)

// TestSearch tests the Search method of the PostgreSQLRepository, that queries the database. It makes use of sqlmock,
// to avoid connecting to a real database when running the tests.
func TestSearch(t *testing.T) {

	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	pgRepo := PostgreSQLRepository{
		conn: conn,
	}
	chapter := domain.Chapter{
		Query: "test",
	}

	rows := sqlmock.NewRows([]string{"id", "title", "chapter"})

	mock.ExpectQuery("SELECT").
		WithArgs(chapter.Query).WillReturnRows(rows)

	_, err = pgRepo.Search(chapter)
	if err != nil {
		t.Error(err)
	}

}

func TestSearch_MultipleQueryFields(t *testing.T) {

	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	pgRepo := PostgreSQLRepository{
		conn: conn,
	}
	chapter := domain.Chapter{
		Query: "arg1 arg2 arg3",
	}

	rows := sqlmock.NewRows([]string{"id", "title", "chapter"})

	mock.ExpectQuery("SELECT").
		WithArgs(chapter.Query).WillReturnRows(rows)

	_, err = pgRepo.Search(chapter)
	if err != nil {
		t.Error(err)
	}

}
