.PHONY: deps clean build

deps:
	go get github.com/aws/aws-lambda-go/lambda

test:
	go test ./... -v
	
clean: 
	rm -rf ./src/main
	
build:
	GOOS=linux GOARCH=amd64 go build -o api/main ./api 

start:
	make build && sam local start-api --port 8080 \
		--region ap-southeast-1 \
		--parameter-overrides "ParameterKey=EnvType,ParameterValue=develop" \

deploy:
	sam deploy -g