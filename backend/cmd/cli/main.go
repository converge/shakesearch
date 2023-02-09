package main

import (
	"bufio"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
	"os"
	"pulley.com/shakesearch/pkg"
)

func main() {

	// prettify logs
	pkg.SetupZeroLog()

	currentPath, err := os.Getwd()
	if err != nil {
		log.Error().Err(err).Msg("failed to get current path")
		return
	}

	// opens eBook reference file
	eBookFile, err := os.OpenFile(currentPath+"/data/pg100-images.html", os.O_RDWR, 0644)
	if err != nil {
		log.Error().Err(err).Msg("failed to open eBookFile")
		return
	}

	doc, err := goquery.NewDocumentFromReader(eBookFile)
	if err != nil {
		log.Error().Err(err).Msg("failed to open eBookFile")
		return
	}
	log.Info().Msg("eBookFile opened successfully")

	// collect chapters
	chapters, err := pkg.GetChapters(doc)
	if err != nil {
		log.Error().Err(err).Msg("failed to get chapters")
		return
	}

	// sqlStatements will contain SQL statements, and will be written to sqlFile
	sqlStatements := &bytes.Buffer{}
	sqlStatements = pkg.CreateInsertStatements(sqlStatements, chapters)
	log.Info().Msg("INSERT SQL statements created successfully")

	// this file was created by migrate tool and is empty in the repository
	sqlFile, err := os.OpenFile(currentPath+"/db/migrations/000002_shakesearch_data.up.sql", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Error().Err(err).Msg("failed to open sqlFile")
		return
	}
	log.Info().Msgf("sqlFile %q opened successfully", sqlFile.Name())

	defer func() {
		err := sqlFile.Close()
		if err != nil {
			log.Error().Err(err).Msg("failed to close sqlFile")
			return
		}
		log.Info().Msg("sqlFile closed successfully")
	}()

	// collect chapters content, those chapters are the main eBook content
	chaptersContent := pkg.GetChapterContent(doc)
	if len(chaptersContent) > 0 {
		log.Info().Msgf("%d chapters content found", len(chaptersContent))
	}

	sqlStatements = pkg.CreateUpdateStatements(sqlStatements, chaptersContent)
	log.Info().Msg("UPDATE SQL statements created successfully")

	sqlFileWriter := bufio.NewWriter(sqlFile)
	if err = sqlFileWriter.Flush(); err != nil {
		return
	}

	_, err = sqlFileWriter.Write(sqlStatements.Bytes())
	if err != nil {
		return
	}
	log.Info().Msg("SQL statements written successfully")
}
