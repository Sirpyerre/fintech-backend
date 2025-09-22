FROM golang:1.24-alpine3.22 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o fintech-backend ./cmd/api

# run stage
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=build /app/fintech-backend .

EXPOSE 8000
CMD ["./fintech-backend"]