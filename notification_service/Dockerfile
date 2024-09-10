FROM golang:1.22.1 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o notification_delivery ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/notification_delivery .

COPY .env .
EXPOSE 6061
CMD ["./notification_delivery"]
