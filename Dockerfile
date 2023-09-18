# Use a smaller base image for the final image
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code (and any necessary files) into the container at /app
COPY . /app

# Build the Go application
RUN go build cmd/main.go


# Use a minimal base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage into the final image
COPY --from=builder /app/main /app/main

# Expose port 8080 for the application
EXPOSE 8080

# Run the application when the container starts
CMD ["./main"]
