# Use the official Golang image as a base
FROM golang:1.20-alpine

# Install curl
RUN apk --no-cache add curl

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o /go-api

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD [ "/go-api" ]