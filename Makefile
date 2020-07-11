.PHONY: deps clean build

deps:

test:
	go test ./... -v
	
clean: 
	rm -rf ./src/main
	
build:
	GOOS=linux GOARCH=amd64 go build -o api/main ./api 

deploy:
	sam deploy -g