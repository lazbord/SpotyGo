FROM golang:1.23-alpine AS auth_builder

WORKDIR /app

COPY . .

WORKDIR /app/services/auth

RUN go mod tidy
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o auth_service .

FROM alpine:latest

WORKDIR /app

COPY --from=auth_builder /app/services/auth/auth_service .

ENV PORT=5000

ENTRYPOINT ["./auth_service"]
