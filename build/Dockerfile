# Start from the latest golang base image
FROM golang:1.20-alpine

# Add Maintainer Info
LABEL maintainer="Vlasov Artem <vlasovartem21@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./internal ./internal
COPY ./cmd ./cmd

# Build the Go app
RUN go build -o main ./cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]