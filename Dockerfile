FROM golang:alpine

ENV GIN_MODE=release

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]