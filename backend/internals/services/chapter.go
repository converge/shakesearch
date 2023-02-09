package services

import "pulley.com/shakesearch/internals/domain"

type SearcherService interface {
	Search(chapter domain.Chapter) ([]domain.Chapter, error)
	GetChapterById(id int64) (domain.Chapter, error)
}

type ChapterService struct {
	Store SearcherService
}

// NewChapterService returns a new ChapterService
func NewChapterService(searcher SearcherService) *ChapterService {
	return &ChapterService{
		Store: searcher,
	}
}

// Search returns a list of search results
func (s *ChapterService) Search(chapter domain.Chapter) ([]domain.Chapter, error) {
	return s.Store.Search(chapter)
}

func (s *ChapterService) GetChapterById(id int64) (domain.Chapter, error) {
	return s.Store.GetChapterById(id)
}
