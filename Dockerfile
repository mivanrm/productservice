# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code (and any necessary files) into the container at /app
COPY . /app


# Build the Go application
RUN go build cmd/main.go

# Expose port 8080 for the application
EXPOSE 8080

# Run the application when the container starts
CMD ["./main"]

