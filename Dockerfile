# Stage 1: Build the Go application
FROM golang:1.22 AS builder
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the application
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/app todo/cmd/*.go

# Stage 2: Create a small image with the built binary
FROM alpine:3.15
WORKDIR /app

ARG IMAGE_TAG
ENV APP_VERSION=$IMAGE_TAG

# Copy the built binary from the builder stage
COPY --from=builder /app/app /app/app

# Copy the .env file
COPY .env ../../.env

# Make sure the binary has execution permissions
RUN chmod +x /app/app

# Set the entrypoint to the built binary
CMD ["/app/app"]
