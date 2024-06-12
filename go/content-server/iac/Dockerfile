FROM alpine:3.19 AS publisher

RUN apk add --no-cache tzdata

WORKDIR /app
COPY bin/content-server ./content

RUN adduser -D appuser
RUN chown -R appuser:appuser /app
USER appuser

EXPOSE 50051

CMD ["/app/content"]


