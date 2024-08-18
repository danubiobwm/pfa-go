# Use the official Golang image as a base
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port the app runs on (if applicable, e.g., for a web server)
EXPOSE 8080

# Run the executable
CMD ["./main"]
