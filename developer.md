# Developer Guide

This guide is intended for developers who want to run the project locally or contribute.

## How to run locally

### Backend

```bash
# start postgres
docker run --name postgresql -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pass123 -p 5432:5432 -d postgres
export DATABASE_URL=postgresql://root:pass123@localhost:5432/postgres?sslmode=disable

cd backend
task run-migration-and-api-start
```

### Frontend

```bash
cd frontend
export API_URL=http://localhost:7001
yarn install
yarn start
```

Open the browser at http://localhost:1234

### Backend tests

Unit tests can be run using task, from the backend folder.

```bash
task unit-test
```

To run the unit tests in a isolation, run the following command:

```bash
docker compose -f docker-compose.yaml up backend-unit-tests
```

End-to-end tests can be run with `task e2e-test` from the backend folder.


### Migrations

Migrations are managed with [golang-migrate], more information can be found in the 
[migration documentation](backend/db/README.md).
