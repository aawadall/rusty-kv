# Docker File for Simple-KV Node Server
# Build Image is Go 1.20.2
FROM golang:1.20.2 AS build

# Set Source Directory
WORKDIR /src

# Copy Source Code
COPY . .

# Get Dependencies
RUN go mod download

# Build Binary
RUN go build -o /bin/simple-kv

# Run Image is Alpine 3.14.0
FROM alpine:3.14.0

# Copy Binary
COPY --from=build /bin/simple-kv /bin/simple-kv

# Set Binary as Entrypoint
ENTRYPOINT ["/bin/simple-kv"]