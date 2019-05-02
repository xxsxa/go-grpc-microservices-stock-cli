build:
	GOOS=linux GOARCH=amd64 go build -o cli.go
	docker build -t shippy-cli-stock .
run:
	docker run shippy-cli-stock