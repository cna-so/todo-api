build :
	@go build -C ./cmd -o ./bin/api
run: build 
	@bin/api
test:
	@go test -v ./...
