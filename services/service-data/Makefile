dependencies:
	go mod download
	go mod tidy

tests: dependencies
	go test ./...

build-local-image:
	docker build . -t airtable-service-data:latest