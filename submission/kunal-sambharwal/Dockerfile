# ---------- Stage 1 : Build ----------
FROM golang:1.26.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o config-service ./cmd

# ---------- Stage 2 : Runtime ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/config-service .

EXPOSE 8080

CMD ["./config-service"]