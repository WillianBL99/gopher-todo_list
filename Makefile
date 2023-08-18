server:
	go run ./cmd/todolist

test-all:
	go test ./pkg/application/usecase/...