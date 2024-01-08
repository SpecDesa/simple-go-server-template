# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# ENV GO111MODULE=on CGO_ENABLED=0

COPY go.mod go.sum /app/
RUN go mod download

# Copy the Go application files into the container at /app
COPY . .

# Build the Go application
RUN go build -o main main.go


# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
