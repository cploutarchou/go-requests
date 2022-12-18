test:
	go test ./...

make test-coverage:
	# test with coverage and output to coverage.html and open it
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out