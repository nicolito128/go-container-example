build:
	go build -o bin/app .

run:
	make build
	./bin/app
