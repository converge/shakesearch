package domain

import (
	"reflect"
	"testing"
)

// TestNewChapter validates the creation of a new Chapter entity, and the validation of the query field.
func TestNewChapter(t *testing.T) {

	tests := []struct {
		name    string
		query   string
		want    *Chapter
		wantErr bool
	}{
		{
			name:  "success NewChapter creation",
			query: "test",
			want: &Chapter{
				Query: "test",
			},
			wantErr: false,
		},
		{
			name:  "fail if query is empty",
			query: "",
			want: &Chapter{
				Query: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewChapter(tt.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewChapter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChapter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
