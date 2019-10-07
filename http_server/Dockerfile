FROM golang:1.11

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN (cd http_server ;go build -o main .)

EXPOSE 8000

CMD ["./http_server/main"]