# syntax=docker/dockerfile:1

# ---- builder ----
FROM golang:1.25-bookworm AS builder

WORKDIR /usr/src/app

# copy module files first for cache
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY . .

# build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" \
    -o app ./cmd/api

# ---- runtime ----
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /usr/src/app/app /app/app

EXPOSE 8080

CMD ["/app/app"]
