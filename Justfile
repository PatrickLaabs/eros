# Copyright © 2024 Patrick Laabs patrick.laabs@me.com

GO_CMD := "go"
TEMPL_CMD := "templ"

update:
    @echo "updating all go.mod debs..."
    {{GO_CMD}} get -u ./...

# Task zum Testen der Go-Anwendung
test:
    @echo "Running tests..."
    {{GO_CMD}} test ./...

cover:
    @echo "Running tests with coverage..."
    {{GO_CMD}} test ./... --cover

cover-html:
    @echo "Creating coverage files and coverage.html"
    {{GO_CMD}} test ./... --coverprofile=./coverage/coverage.out
    {{GO_CMD}} tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html

bench:
    @echo "Running benchmarks.."
    {{GO_CMD}} test -bench=./...

# Task zur Formatierung des Go-Codes
fmt:
    @echo "Formatting the code..."
    {{GO_CMD}} fmt ./...

# Task zum Linting des Go-Codes
lint:
    @echo "Linting the code..."
    {{GO_CMD}} vet ./...

gen:
    @echo "Generating templ files..."
    {{TEMPL_CMD}} generate ./frontend/views

swagger:
    @echo "Generating API documentation..."
    swag init -g ./api/start.go -o ./docs

build: server frontend

server:
    @echo "Building API Server Binary..."
    {{GO_CMD}} build ./cmd/api-server/api-server.go

frontend:
    @echo "Building Frontend Server Binary..."
    {{GO_CMD}} build ./cmd/frontend-server/frontend-server.go
