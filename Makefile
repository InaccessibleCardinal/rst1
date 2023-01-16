build:
	go build -o bin/rst1

run:
	go run ./app/main.go

test:
	go test ./...