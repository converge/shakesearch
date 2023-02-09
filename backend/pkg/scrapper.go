package pkg

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"pulley.com/shakesearch/internals/domain"
	"strconv"
	"strings"
)

// GetChapters uses goquery library to parse Shakespeare eBook chapters, and returns a slice of Chapter. Chapter id is
// used as key to match the chapter content.
func GetChapters(doc *goquery.Document) ([]domain.Chapter, error) {

	var chapters []domain.Chapter
	doc.Find("div:nth-child(7) table:nth-child(2) tbody tr td a").Each(func(i int, s *goquery.Selection) {
		var c domain.Chapter
		chapterId, _ := s.Attr("href")
		// todo: check errors
		// remove chap
		convertedId := strings.Replace(chapterId, "#chap", "", -1)
		c.Id, _ = strconv.ParseInt(convertedId, 10, 64)
		c.Title = s.Text()

		chapters = append(chapters, c)
	})

	return chapters, nil
}

// GetChapterContent uses goquery library to parse Shakespeare eBook chapters content, and returns a slice of Chapter
// Content. The content provided will be used to update the chapter content in the database.
func GetChapterContent(doc *goquery.Document) []domain.Chapter {

	var chaptersContent []domain.Chapter

	doc.Find("body > div.chapter").Each(func(i int, s *goquery.Selection) {

		var chapter domain.Chapter

		// find the chapter link via css selector h2 a
		chapterLink := s.Find("h2 a")

		// get link attribute id
		chapterId, _ := chapterLink.Attr("id")

		// there are different kind of chapters, we only want the ones with chap prefix, that matches the HTML pattern
		// for real chapters content
		if strings.HasPrefix(chapterId, "chap") {

			// remove chap, make it available for int casting, so we can use it as key
			convertedId := strings.Replace(chapterId, "chap", "", -1)
			// todo: handle error
			id, _ := strconv.ParseInt(convertedId, 10, 64)

			escapedString := strings.ReplaceAll(s.Text(), `'`, `''`)

			chapter.Id = id
			chapter.Content = escapedString

			// appending
			chaptersContent = append(chaptersContent, chapter)

		}
	})

	return chaptersContent
}

// CreateInsertStatements creates SQL insert statements for the chapters slice, and returns a buffer with the content.
func CreateInsertStatements(content *bytes.Buffer, chapters []domain.Chapter) *bytes.Buffer {

	for _, c := range chapters {
		content.WriteString(fmt.Sprintf("INSERT INTO shakesearch (id, title, created) VALUES (%d, '%s', NOW());\n", c.Id, c.Title))
	}

	return content
}

// CreateUpdateStatements creates SQL update statements for the chapters slice, and returns a buffer with the content.
func CreateUpdateStatements(content *bytes.Buffer, chapters []domain.Chapter) *bytes.Buffer {

	for _, c := range chapters {
		content.WriteString(fmt.Sprintf("UPDATE shakesearch SET chapter = '%s' WHERE id = %d;\n", c.Content, c.Id))
	}

	return content
}
