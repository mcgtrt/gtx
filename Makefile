# Build, run, and test project
build: 
	@templ generate
	@go build -o bin/sendiforce

run: build
	@./bin/sendiforce

test:
	@go test -v ./... -count=1 -cover

# Build and watch
templ:
	@templ generate --watch --proxy=http://localhost:3069

css:
	@npx tailwindcss -i ./view/css/style.css -o ./public/style.css --watch

# Kill local process on port :3069
kill:
	@lsof -t -i tcp:3069 | xargs kill