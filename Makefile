server:
	go run ./cmd/todolist

test-all:
	go test ./internal/application/usecase/...