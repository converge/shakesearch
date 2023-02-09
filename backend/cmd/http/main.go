package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"pulley.com/shakesearch/internals/handlers"
	"pulley.com/shakesearch/internals/repositories"
	"pulley.com/shakesearch/internals/services"
	"pulley.com/shakesearch/pkg"
)

func main() {

	// logging setup
	pkg.SetupZeroLog()

	databaseURL := os.Getenv("DATABASE_URL")
	if len(databaseURL) == 0 {
		log.Panic().Msg("DATABASE_URL is not set")
	}

	// instantiate database use pgx driver for PostgreSQL
	conn, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Error().Err(err).Msg("unable to open database connection")
		return
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Msg("unable to close database connection")
		}
	}(conn)

	// loose coupled components setup to enable dependency injection
	// constructors expects an interface to be satisfied, in this way, the components can be easily replaced by another
	// implementation
	postgreSQLRepository := repositories.NewPostgresRepository(conn)
	chapterService := services.NewChapterService(postgreSQLRepository)
	chapterHandler := handlers.NewChapterHandler(chapterService)

	// Mux router provides nice routing management
	router := mux.NewRouter()
	router.StrictSlash(true)
	subRouter := router.PathPrefix("/v1").Subrouter()

	// serve swagger UI
	swaggerRouter := http.StripPrefix("/v1/swagger-ui", http.FileServer(http.Dir("./third_party/swagger-ui/")))
	subRouter.PathPrefix("/swagger-ui").Handler(swaggerRouter)

	// CORS middleware
	subRouter.Use(pkg.CorsMiddleware)

	// application endpoints
	subRouter.HandleFunc("/", handlers.Version).Methods(http.MethodGet)
	subRouter.HandleFunc("/chapter/search", chapterHandler.Search).Methods(http.MethodGet)
	subRouter.HandleFunc("/chapter/{id}", chapterHandler.GetChapterById).Methods(http.MethodGet)

	port := ":7001"
	log.Info().Msgf("starting http server at localhost%s", port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Error().Err(err).Msg("failed to start http server")
	}
}
