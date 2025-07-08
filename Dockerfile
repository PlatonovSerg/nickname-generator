FROM golang:1.24.1-alpine AS builder
WORKDIR /app
RUN apk add --no-cache gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=1
RUN go build -o nickname-generator ./cmd/main.go

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates gcc musl-dev
COPY --from=builder /app/nickname-generator .
COPY data/words.db ./data/words.db
EXPOSE 8080
ENV GIN_MODE=release
CMD ["./nickname-generator"]