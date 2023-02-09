package services

import (
	"pulley.com/shakesearch/internals/domain"
	"reflect"
	"testing"
)

type mockChapterService struct{}

func (a mockChapterService) GetChapterById(id int64) (domain.Chapter, error) {
	return domain.Chapter{}, nil
}

func (a mockChapterService) Search(chapter domain.Chapter) ([]domain.Chapter, error) {
	return nil, nil
}

// TestNewChapterService tests the creation of a new ChapterService.
func TestNewChapterService(t *testing.T) {

	mockService := mockChapterService{}

	tests := []struct {
		name     string
		searcher SearcherService
		want     *ChapterService
	}{
		{
			name:     "TestNewChapterService",
			searcher: nil,
			want: &ChapterService{
				Store: nil,
			},
		},
		{
			name:     "TestNewChapterService",
			searcher: mockService,
			want: &ChapterService{
				Store: mockService,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChapterService(tt.searcher); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChapterService() = %v, want %v", got, tt.want)
			}
		})
	}
}
