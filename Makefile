mock:
	mockery --all --case underscore

test:
	go test -race -coverprofile=coverage.out $(shell go list ./... | grep -v /mocks/ | grep -v /cmd/)
	go tool cover -func=coverage.out
	rm coverage.out

lint:
	golang-lint run ./...