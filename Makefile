clean:
	rm -rf build

build: clean
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/build ./cmd/main.go

scp: build
	scp build/build main:/home/arvindpunk/word-proximity-service/