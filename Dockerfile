FROM golang:1.23 as builder

# Maintainer info
LABEL maintainer="Arafate do Rosario"

# # Intall git
# # Git is required for fetching the dependencies
# RUN apk update && apk add --no-cacahe git && apk add --no-cache bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .
COPY .env .

# # Download all the dependencies
# RUN go get -d -v ./...

# # Install the package
# RUN go install -v ./...

# Build the Go app
# RUN go build -o /build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build

# FROM alpine
# RUN apk add --no-cache ca-certificates && update-ca-certificates
EXPOSE 8080 8080


# # Expose port 8080 to the internet
# EXPOSE 8080

# Run the executable
CMD ["/build"]