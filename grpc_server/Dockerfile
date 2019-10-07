FROM golang:1.11

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN (cd grpc_server ;go build -o main .)

EXPOSE 9000

CMD ["./grpc_server/main"]