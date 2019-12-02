FROM golang

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN go build cmd/paytest-server/main.go

CMD ["./main"]
