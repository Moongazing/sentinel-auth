# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o sentinel cmd/main.go

# Stage 2: Run
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/sentinel .
COPY .env .

EXPOSE 8080

CMD ["./sentinel"]
