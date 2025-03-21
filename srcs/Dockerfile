# Build Stage
FROM golang:1.24.1 AS builder
WORKDIR /srv

# Copy and download dependencies
COPY go.mod ./
RUN go mod tidy

# Copy source code and build
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o server .
RUN go build -o server .

# Final Stage (Slim Image)
FROM alpine:latest
ENV PATH="/srv:${PATH}"
WORKDIR /srv

# Install CA certificates (required for some HTTPS connections)
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache libc6-compat

# Copy built binary from builder
COPY --from=builder /srv/server .
COPY index.html /srv/
COPY assets/ /srv/assets

# Ensure the binary is executable
RUN chmod +x /srv/server

# Expose the srvlication port
EXPOSE 8080

# Run the binary
CMD ["./server"]
