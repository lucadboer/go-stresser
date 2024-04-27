FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o cli-stress-test .

FROM alpine

WORKDIR /app

COPY --from=builder /app/cli-stress-test .

ENTRYPOINT ["./cli-stress-test"]
