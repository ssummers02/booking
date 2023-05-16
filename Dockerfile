# Stage 1: Build the application
FROM golang:1.20-alpine AS builder

ENV GOPATH=/
RUN go version

# Install psql and other dependencies
RUN apk update
RUN apk add --no-cache postgresql-client git

# Copy the source code
COPY ./ ./

# Download dependencies and build the application
RUN go mod download
RUN go build -o /app ./cmd/main.go

# Stage 2: Create the final image
FROM alpine:latest

# Install psql
RUN apk update
RUN apk add --no-cache postgresql-client

# Copy the compiled application from the builder stage
COPY --from=builder /app /app

# Expose the application port
EXPOSE 8081 8081

# Set the entrypoint
ENTRYPOINT ["./app"]