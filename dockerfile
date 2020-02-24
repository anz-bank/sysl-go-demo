# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM alpine:latest

# Add Maintainer Info
LABEL maintainer="Joshua Carpeggiani <joshua.carpeggiani@.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY main ./

# Expose port 8080 to the outside world
EXPOSE 80
EXPOSE 8080

# Command to run the executable
CMD ["./main"]