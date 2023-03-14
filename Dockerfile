FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start fresh from a smaller image
FROM alpine:latest

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Command to run the executable
CMD ["/app/main"]
