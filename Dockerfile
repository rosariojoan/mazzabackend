# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder

# Maintainer info
LABEL maintainer="Arafate do Rosario"

# Setup folders
WORKDIR /app

# Set environment variables
ENV DB_DRIVER=${DB_DRIVER} \
    DB_USER=${DB_USER} \
    DB_PASSWORD=${DB_PASSWORD} \
    DB_NAME=${DB_NAME} \
    DB_HOST=${DB_HOST} \
    DB_PORT=${DB_PORT} \
    DB_URL=${DB_URL} \
    PORT=${PORT} \
    SECRET_KEY=${SECRET_KEY} \
    ACCESS_TOKEN_EXPIRE=${ACCESS_TOKEN_EXPIRE} \
    FIREBASE_CONFIG_FILE=${FIREBASE_CONFIG_FILE} \
    FIREBASE_BUCKET=${FIREBASE_BUCKET}

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/build ./app

# Debugging: List contents of /app
# RUN ls -la /app

# Use a minimal alpine image for the final stage
FROM alpine:latest

# Install ca-certificates
# RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/build .

# Copy the template files
COPY --from=builder /app/app/utils/email/templates /app/app/utils/email/templates

# Copy assets
COPY --from=builder /app/app/assets /app/app/assets

# Ensure the binary is executable
RUN chmod +x /app/build

# Expose port 8080 to the internet
EXPOSE 8080

# Run the executable
CMD ["/app/build"]