tidy:
	go mod tidy
	go fmt ./...

run: tidy
	go run ./...

build: tidy
	go build -o ./dist/go-webforms ./... 

clean:
	rm -rf ./dist