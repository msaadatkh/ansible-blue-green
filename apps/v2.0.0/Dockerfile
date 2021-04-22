# TODO v2.0.0
# Builder stage
FROM golang:1.15-alpine as builder

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod download
RUN go mod verify

# Build the Go app
RUN go build -o go-app

# Final stage
FROM golang:1.15-alpine

RUN apk add --no-cache curl

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy Go app from builder stage
COPY --from=builder /app/go-app .

HEALTHCHECK --interval=5s --timeout=2s --retries=3 \
  CMD curl -f http://localhost:3000/healthz || exit 1

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the binary program produced by `go install`
CMD ["./go-app"]