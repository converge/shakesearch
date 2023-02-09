package domain

import (
	"errors"
	"strings"
)

type Chapter struct {
	Query   string `json:"query,omitempty"`
	Id      int64  `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

// NewChapter creates a new Chapter, that is a domain specific entity that represents a eBook chapter. During the
// creation of the Chapter, the query is validated to be non-empty, and if multiple query entries are provided, they are
// joined with the AND operator to be later consumed by full-text search.
func NewChapter(query string) (*Chapter, error) {

	if len(query) == 0 {
		return &Chapter{}, errors.New("query is empty")
	}

	query = strings.Replace(query, " ", " & ", -1)

	return &Chapter{
		Query: query,
	}, nil
}
