FROM golang:1.15 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gogoat .

# Start a new stage from scratch for the runtime
FROM alpine:latest  

# Set the working directory
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/gogoat .

# Add ca-certificates in case your application makes external HTTPS requests
RUN apk --no-cache add ca-certificates

# Expose HTTP port
EXPOSE 8080

# Command to run the executable
CMD ["./gogoat"]
