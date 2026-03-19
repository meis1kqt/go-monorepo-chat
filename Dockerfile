FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./service/auth/cmd/main.go

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/auth-service .
COPY config/auth-local.yaml config/auth-local.yaml
EXPOSE 50051
