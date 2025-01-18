# syntax=docker/dockerfile:1

# Use the official Golang image as a build stage
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o r314dash .

# Use a Node.js image to build Tailwind CSS
FROM node:16 AS tailwind-builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json (if available)
COPY package.json package-lock.json ./

# Install npm dependencies
RUN npm install

# Copy the source code into the container
COPY . .

# Build Tailwind CSS
RUN npm run build

# Start a new stage from scratch
FROM alpine:latest

#ENV PORT=8080

# Set the Current Working Directory inside the container
WORKDIR /app

# Create a non-root user with UID 1000
RUN addgroup -S r314dash && adduser -S r314dash -u 1000 -G r314dash && apk add curl

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/r314dash /app/config.yaml ./
# Copy the built Tailwind CSS files from the tailwind-builder stage
COPY --from=tailwind-builder /app/static ./static
COPY --from=tailwind-builder /app/templates ./templates
#COPY entrypoint.sh .

# Make the entrypoint script executable
#RUN chmod +x ./entrypoint.sh

EXPOSE 3000

# Change ownership of the application files to the non-root user
RUN chown -R r314dash:r314dash /app

# Switch to the non-root user
USER r314dash

# Command to run the executable
CMD ["./r314dash"]
#CMD ["entrypoint.sh"]
