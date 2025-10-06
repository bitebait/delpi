FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN addgroup -g 1000 appuser && adduser -D -u 1000 -G appuser appuser

WORKDIR /home/appuser
COPY --from=builder /app/server .
COPY .env .

RUN chown appuser:appuser server .env

USER appuser

CMD ["./server"]
