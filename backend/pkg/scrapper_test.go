package pkg

import (
	"bytes"
	"pulley.com/shakesearch/internals/domain"
	"reflect"
	"testing"
)

func TestCreateInsertStatements(t *testing.T) {

	tests := []struct {
		name     string
		content  *bytes.Buffer
		chapters []domain.Chapter
		want     *bytes.Buffer
	}{
		{
			name:     "test empty chapters slice",
			content:  bytes.NewBufferString(""),
			chapters: []domain.Chapter{},
			want:     bytes.NewBufferString(""),
		},
		{
			name:    "success on creating insert statements",
			content: bytes.NewBufferString(""),
			chapters: []domain.Chapter{
				{
					Id:    1,
					Title: "Test",
				},
			},
			want: bytes.NewBufferString("INSERT INTO shakesearch (id, title, created) VALUES (1, 'Test', NOW());\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateInsertStatements(tt.content, tt.chapters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInsertStatements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUpdateStatements(t *testing.T) {

	tests := []struct {
		name     string
		content  *bytes.Buffer
		chapters []domain.Chapter
		want     *bytes.Buffer
	}{
		{
			name:     "test empty chapters slice",
			content:  bytes.NewBufferString(""),
			chapters: []domain.Chapter{},
			want:     bytes.NewBufferString(""),
		},
		{
			name:    "success on creating update statements",
			content: bytes.NewBufferString(""),
			chapters: []domain.Chapter{
				{
					Id:      1,
					Content: "Test",
				},
			},
			want: bytes.NewBufferString("UPDATE shakesearch SET chapter = 'Test' WHERE id = 1;\n"),
		},
		{
			name:    "success on creating update statements multiple chapters",
			content: bytes.NewBufferString(""),
			chapters: []domain.Chapter{
				{
					Id:      1,
					Content: "Test",
				},
				{
					Id:      2,
					Content: "chapter content",
				},
			},
			want: bytes.NewBufferString("UPDATE shakesearch SET chapter = 'Test' WHERE id = 1;\nUPDATE shakesearch SET chapter = 'chapter content' WHERE id = 2;\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateUpdateStatements(tt.content, tt.chapters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUpdateStatements() = %v, want %v", got, tt.want)
			}
		})
	}
}
