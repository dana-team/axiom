FROM golang:1.25.2 AS builder

WORKDIR /workspace

# Copy dependency definitions
COPY backend/go.mod backend/go.sum backend/
COPY operator/go.mod operator/go.sum operator/

# Download dependencies
WORKDIR /workspace/backend
RUN go mod download

# Copy source
WORKDIR /workspace
COPY backend/ backend/
COPY operator/ operator/

# Build
WORKDIR /workspace/backend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /workspace/backend/server /app/server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/server"]
