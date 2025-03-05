FROM golang:alpine3.21 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o server .

FROM alpine:3.21

WORKDIR /root

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]