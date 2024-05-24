FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o loadtester .

FROM alpine

WORKDIR /app

COPY --from=builder /app/loadtester .

ENTRYPOINT ["./loadtester"]
