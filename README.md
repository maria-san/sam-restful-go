# sam-restful-server

AWS SAM RESTful server

## Tech
- AWS SAM Framework
- REST
- Go
- Serverless

## Installation
### Local
    make start
### Deploy
    make deploy
        

## API Resources

  - [GET /hello](#get-hello)
  - [POST /hi](#post-hi)

### GET /hello

Example: {{local}}/{version}/hello

Response body:

    "Hello"


### POST /hi

Example: {{local}}/{version}/hello

Request body:

    {
        "name": "maria",
    }

Response body:

    "Hi maria!"
