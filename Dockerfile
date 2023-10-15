# Use an official Go image as a base
FROM golang:1.19

# Install Redis from the official repository
RUN apt-get update && apt-get install -y redis-server

# Copy your Go project into the container
WORKDIR /app
COPY . .

# Build your Go application
RUN go build -o myapp .

# Expose ports
EXPOSE 6379 8080

# Set environment variables
ENV REDIS_HOST=redis
ENV REDIS_PORT=6379

# Start Redis and your application using a shell script
CMD ["./myapp"]
