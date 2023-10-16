# Use an official Go image as a base
FROM golang:1.19

# Copy your Go project into the container
WORKDIR /app
COPY . .

# Build your Go application
RUN go build -o myapp

# Expose port
EXPOSE 8080

# Start your application
CMD ["./myapp"]
