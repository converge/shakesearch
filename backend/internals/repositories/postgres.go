package repositories

import (
	"database/sql"
	"pulley.com/shakesearch/internals/domain"
)

type PostgreSQLRepository struct {
	conn *sql.DB
}

// NewPostgresRepository is a constructor for PostgreSQLRepository, it expects a database connection.
func NewPostgresRepository(conn *sql.DB) *PostgreSQLRepository {
	return &PostgreSQLRepository{
		conn: conn,
	}
}

// Search is a method that returns a list of search results based on the Chapter query. It uses PostgreSQL full text
// search feature.
func (r *PostgreSQLRepository) Search(chapter domain.Chapter) ([]domain.Chapter, error) {

	// full text search query on title and chapter. The order by clause is used to order the results that match the
	// title over the ones that match the chapter. The eBook content is english only, so the vector setup language is
	// set to english.

	// to_tsvector converts the text to a token vector, and will also remove words like 'the', 'and', 'on', on which we
	// don't want to search for. lexeme words are also considered, ex. love will match loving, loved, etc.
	sql := `SELECT id, title, LEFT(chapter, 300) as chapter
			  FROM shakesearch
			 WHERE to_tsvector('english', title || ' ' || chapter) @@ plainto_tsquery('english', $1)
			 ORDER BY 
			     CASE
					WHEN (title @@ plainto_tsquery('english', $1)) THEN 0
					ELSE 1
				 END DESC;`

	rows, err := r.conn.Query(sql, chapter.Query)
	if err != nil {
		return []domain.Chapter{}, err
	}

	var results []domain.Chapter

	for rows.Next() {
		result := domain.Chapter{}
		if err = rows.Scan(&result.Id, &result.Title, &result.Content); err != nil {
			return []domain.Chapter{}, err
		}
		results = append(results, result)
	}

	return results, nil
}

// GetChapterById returns a single chapter based on the id.
func (r *PostgreSQLRepository) GetChapterById(id int64) (domain.Chapter, error) {
	sql := `SELECT id, title, chapter
			  FROM shakesearch
			 WHERE id = $1;`

	result := domain.Chapter{}
	if err := r.conn.QueryRow(sql, id).Scan(&result.Id, &result.Title, &result.Content); err != nil {
		return domain.Chapter{}, err
	}

	return result, nil
}
