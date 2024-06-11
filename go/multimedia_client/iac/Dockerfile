FROM golang:alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY config/ ./config/
COPY internal ./internal/

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/client ./cmd/main.go

FROM alpine:3.19 AS publisher

RUN apk add --no-cache tzdata

WORKDIR /app
COPY --from=builder /app/bin/client ./client
COPY public ./public/

RUN adduser -D appuser
RUN chown -R appuser:appuser /app
USER appuser

EXPOSE 3000

CMD ["/app/client"]


