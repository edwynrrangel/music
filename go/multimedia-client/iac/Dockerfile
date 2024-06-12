FROM alpine:3.19 AS publisher

RUN apk add --no-cache tzdata

WORKDIR /app
COPY bin/multimedia-client ./client
COPY public ./public/

RUN adduser -D appuser
RUN chown -R appuser:appuser /app
USER appuser

EXPOSE 3000

CMD ["/app/client"]


