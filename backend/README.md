# Shakesearch backend

## Backend structure

The backend uses Hexagonal Architecture to enable separation of concerns, loose coupled components, and easy testing. It
also uses the concept of DDD(domain driven design) to separate the domain logic from the rest of the application.

Project tree explained:

```bash
├── Dockerfile.dev # dockerfile for development
├── README.md
├── Taskfile.yaml # tasks automation
├── cmd # drivers
│   ├── cli # cli tool to convert the eBook content into database entries
│   └── http # http server
├── data # data source (eBook)
│   └── pg100-images.html
├── db
│   ├── README.md
│   └── migrations # managed by Go tool migration
├── go.mod
├── go.sum
├── internals
│   ├── domain # domain logic
│   ├── handlers # http handlers
│   ├── repositories # db repositories
│   └── services # services
├── pkg
│   ├── constants.go
│   ├── log.go
│   ├── middlewares.go
│   ├── scrapper.go
│   └── scrapper_test.go
├── tests # e2e test
│   ├── README.md
│   └── search_test.go
└── third_party
    ├── README.md
    └── swagger-ui # swagger ui for api docs
```

## Dependencies

There is a minimum use of direct external dependencies, they are:

```bash
github.com/DATA-DOG/go-sqlmock # for testing
github.com/PuerkitoBio/goquery # for parsing the html
github.com/chromedp/chromedp   # for the end to end test
github.com/gorilla/mux         # for http routing
github.com/jackc/pgx/v5        # for postgres driver
github.com/rs/zerolog          # for logging
```
