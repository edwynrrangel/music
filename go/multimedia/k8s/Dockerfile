FROM golang:alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY config/ ./config/
COPY internal ./internal/

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/multimedia ./cmd/main.go

FROM alpine:3.19 AS publisher

RUN apk add --no-cache tzdata

WORKDIR /app
COPY --from=builder /app/bin/multimedia ./multimedia

RUN adduser -D appuser
RUN chown -R appuser:appuser /app
USER appuser

EXPOSE 50051

CMD ["/app/multimedia"]


