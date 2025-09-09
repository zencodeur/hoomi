# Hoomi Backend

Backend for the Hoomi application built with Go and Gin.

## Structure

```
backend/
├── main.go          # Entry point
├── go.mod           # Go modules
├── go.sum           # Go checksums
├── .env.example     # Environment variables example
├── .gitignore       # Git ignore rules
└── ...
```

## Development

1. Install Go (1.21 or later)
2. Install dependencies: `go mod tidy`
3. Copy `.env.example` to `.env` and configure
4. Run: `go run main.go`

## API

The API will be documented in OpenAPI format in `docs/api.md`.