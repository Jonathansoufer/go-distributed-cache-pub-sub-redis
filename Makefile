build:
	go build -o bin/go-distributed-cache-pub-sub-redis -v

run: build
	./bin/go-distributed-cache-pub-sub-redis