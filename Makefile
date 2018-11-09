validate:
	gometalinter --exclude 'pkg/mod' --deadline=180s ./...

test:
	go test ./... -coverprofile coverage.out
	go tool cover -html=coverage.out -o coverage.html

.PHONY: validate test
