FROM golang:1.20 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o resumme-builder

EXPOSE 9000

FROM ubuntu:latest AS launcher

# Download google-chrome
RUN apt-get update && apt-get install -y wget
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt-get install -y ./google-chrome-stable_current_amd64.deb

COPY --from=builder /app .
CMD ["./resumme-builder", "server"]