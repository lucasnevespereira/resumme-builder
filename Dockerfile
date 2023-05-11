FROM golang:1.20 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o resumme-builder

EXPOSE 9000

FROM alpine:latest AS launcher
COPY --from=builder /app .
CMD ["./resumme-builder", "server"]