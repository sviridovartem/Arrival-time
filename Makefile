build:
	go build ./cmd/backend-core/main.go

run:
	cd scripts && docker-compose up -d
	
test:
	go test ./...