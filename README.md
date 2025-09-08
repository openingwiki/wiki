# wiki
Wiki service. Responsible for storing opening, ending, etc.

## Endpoints

- POST `/api/v1/anime` {"title": string}
- GET `/api/v1/anime/{id}`
- POST `/api/v1/singers` {"name": string}
- GET `/api/v1/singers/{id}`
- POST `/api/v1/openings` {"anime_id": int, "singer_id": int, "type": opening|ending|ost, "title": string, "order_number": int}
- GET `/api/v1/openings/{id}`

## Run

Install Go (1.25.1+) and run:

```bash
go mod tidy
go run ./cmd/server
```

## Make targets

```bash
make tidy   # go mod tidy
make test   # run unit tests with race detector
make lint   # run golangci-lint
make ci     # tidy + lint + test
```
