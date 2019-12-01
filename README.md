# Watchy API

This is the API for the watchy app

## Getting Started

```bash
# clone project
go get github.com/Wuzi/pay-test-go

# run server
go run ./cmd/paytest-server/main.go --scheme http --port 3000

# build server
go build ./cmd/paytest-server
```

## Documentation

Start the server and go to http://localhost:3000/v1/docs

## Generating new documentation

You should always update `swagger.yml` first, then regenerate server with

```bash
go generate ./restapi
```

## Built With

* [go-swagger](https://goswagger.io/) - Swagger 2.0 implementation for go

## Contributing

Contributions are greatly appreciated. Please fork this repository and open a pull request to add snippets, make grammar tweaks, improve project structure, etc.
