FROM golang:1.23-alpine AS files_builder

WORKDIR /app

COPY . .

WORKDIR /app/services/files

RUN go mod tidy
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o files_service .

FROM alpine:latest

WORKDIR /app

RUN apk update && \
    apk add --no-cache ffmpeg

COPY --from=files_builder /app/services/files/files_service .

ENV PORT=5001

ENTRYPOINT ["./files_service"]
