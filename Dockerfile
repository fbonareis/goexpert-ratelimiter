FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/ratelimiter
CMD ["./app"]
