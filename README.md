# ShakeSearch

## Overview

This is a web app that allows a user to search for a text string in the complete works of Shakespeare.

 * [Backend Overview](backend/README.md)
 * [Developer Guide](developer.md)
 * [OpenAPI specification](backend/third_party/swagger-ui/shakesearch.yaml)
 * [End-to-End test](backend/tests/README.md)

### How it works

It loads the complete works of Shakespeare into the database (PostgreSQL), and then allows the user to search for a text
string in the database. The data parsing is done using goquery, and the search is done using PostgreSQL full text 
search.

A CLI tool at [backend/cmd/cli](backend/cmd/cli/main.go) is used to parse the data and create a migration file. The 
migration file is then run to load the data into the database. (the idea is to automate the process during CI/CD).

After data being parsed and loaded into the database, the backend exposes an API to search for the text string. The
search is done using PostgreSQL full text search. And the weight of the title and content(chapter content) were balanced
to get the best result. More details can be found [here](backend/db/migrations/000001_shakesearch_initial.up.sql).

The frontend is a simple React app that allows the user to search for a text string and displays the results.

