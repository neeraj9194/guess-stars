build:
	go build src/cli/game.go

run:
	go run src/cli/game.go

test:
	go test ./... -v