# ---- Build stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git for go modules if needed
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# ---- Runtime stage ----
FROM alpine:3.19

WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata

# Non-root user
RUN adduser -D -g '' appuser
USER appuser

COPY --from=builder /app/server /app/server

EXPOSE 8080
ENTRYPOINT ["/app/server"]
