# link_shortener

Creating a unique short link  
Service provides API with Create and Get methods using grpc  
To use, generate a client from the description in the proto file, or use the command below  

## Build

`docker build -t linkshortener .`  
`docker run -p 8080:8080 linkshortener`

## Usage

`go run tests/test.go`

## Tests

`docker build -t tester -f tests/Dockerfile .`  
`docker run tester`
