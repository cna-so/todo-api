build :
	@go build -C ./cmd -o ../bin/api
run: build 
	@bin/api
clean:
	@go clean -testcache
test:clean
	@go test -v ./...
