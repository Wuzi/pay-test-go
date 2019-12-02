# Pay Test API

This is the API for the pay test

## Getting Started

```bash
# clone project
go get github.com/Wuzi/pay-test-go

# run server
go run ./cmd/paytest-server/main.go --scheme http --port 3002

# build server
go build ./cmd/paytest-server
```

## Documentation

Start the server and go to http://localhost:3002/v1/docs

## Generating new documentation

You should always update `swagger.yml` first, then regenerate server with

```bash
go generate ./restapi
```

## Docker
To run the app with docker first build the image:

```bash
docker build -t wuzi/pay-test-go:0.1.0 .
```

Then start the container:

```bash
docker-compose up
```

And go to http://localhost:3002

## Built With

* [go-swagger](https://goswagger.io/) - Swagger 2.0 implementation for go

## Contributing

Contributions are greatly appreciated. Please fork this repository and open a pull request to add snippets, make grammar tweaks, improve project structure, etc.
