tailwind:
	@./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

tailwind-watch:
	@./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

templ:
	@.bin/templ generate

templ-watch:
	@templ generate --watch

build: tailwind templ
	@go build -o ./bin/gtx ./cmd/main.go

run: build
	@./bin/gtx

test:
	@go test -race -v -timeout 30s ./...