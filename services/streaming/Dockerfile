FROM golang:1.23-alpine AS streaming_builder

WORKDIR /app

COPY . .

WORKDIR /app/services/streaming

RUN go mod tidy
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o streaming_service .

FROM alpine:latest

WORKDIR /app

COPY --from=streaming_builder /app/services/streaming/streaming_service .

ENV PORT=5002

ENTRYPOINT ["./streaming_service"]
