all:
	@go build -o interview
	@./interview

test:
	go test