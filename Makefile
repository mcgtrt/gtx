tailwind:
	@npx tailwindcss -i ./src/css/input.css -o ./src/css/style.min.css --minify

tailwind-watch:
	@npx tailwindcss -i ./src/css/input.css -o ./src/css/style.css --watch

templ:
	@./bin/templ generate

templ-watch:
	@./bin/templ generate --watch

build: 
	@go build -o ./bin/gtx ./cmd/main.go

run: tailwind templ build 
	@./bin/gtx && air

test:
	@go test -race -v -timeout 30s ./...

.PHONY: tailwind