FROM golang:1.23.1-alpine3.19 AS build

# Create group and non-root user to prevent malicious users from accessing root user
RUN addgroup -S blockchaingroup && adduser -S blockuser -G blockchaingroup

# Create working directory
WORKDIR /app

# Change ownership from root user to non-root user initially so permission errors do not arise
RUN chown -R blockuser:blockchaingroup /app

# Switch to root user
USER blockuser

# Copy go.mod and go.sum files into working directory
# Explicit declaration of non-root ownership forces command to run under non-root user to prevent permission errors
COPY --chown=blockuser:blockchaingroup go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy program files into working directory
COPY --chown=blockuser:blockchaingroup . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o main ./main.go

# Runtime stage base image
FROM alpine:3.19

# Add packages
RUN apk --no-cache add ca-certificates

# Create group and non-root user to prevent malicious users from accessing root user
# User created in build stage is removed after 
RUN addgroup -S blockchaingroup && adduser -S blockuser -G blockchaingroup

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=build --chown=blockuser:blockchaingroup /app/main .

# Switch to the non-root user
USER blockuser

# Command to run the application
CMD ["./main"]