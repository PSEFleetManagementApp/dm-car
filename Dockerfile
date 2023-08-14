# ==================
# BUILD STAGE
# ==================
# Use a lightweight Go image based on Alpine for building the application
FROM golang:1.21.0-alpine3.18 AS build

# Set the working directory
WORKDIR /app

# Copy go module files to the container's filesystem
COPY go.mod go.sum ./

# Download the dependencies as specified in the go.mod and go.sum
# To optimize Docker's caching mechanism, download dependencies prior to copying the rest of the source code
RUN go mod download

# Copy the entire source code to the container
COPY . .

# Compile the Go application with necessary flags
# Disable CGO and set the target OS to Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# ==================
# TEST STAGE
# ==================
# Use the previous build stage as a base for test stage
FROM build AS test

# Execute all unit tests
RUN go test -v ./...

# ==================
# PRODUCTION STAGE
# ==================
# Use a minimal Alpine image for the final production container
FROM alpine:3.18 AS production

# Set the working directory in the container to /root/
WORKDIR /root/

# Copy the main binary from the build stage to the production image
COPY --from=build /app/main .

# Inform Docker that the container listens on port 80
EXPOSE 80

# Execute the main binary at startup
CMD ["./main"]