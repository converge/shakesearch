# End-to-end tests

This folder contains end-to-end tests for the backend.

### How to run the end-to-end tests:

**1. Start the backend server locally**

```bash
cd backend
go run cmd/http/main.go
```

**2. Start the frontend server locally**

```bash
cd frontend
yarn start
```

**3. Run the e2e tests**

```bash
cd backend
task e2e-tests
```

note: this folder contains end-to-end tests, unit tests are available close to the code they test.
