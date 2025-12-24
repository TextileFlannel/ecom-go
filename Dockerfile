FROM golang:1.24.10-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o todo-app ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/todo-app .

EXPOSE 8080

CMD ["./todo-app"]