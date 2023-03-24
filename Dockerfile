# Docker File for Simple-KV Node Server
# Build Image is Go 1.20.2
FROM golang:1.20.2-alpine3.17 AS build

# Set Source Directory
WORKDIR /src

# Copy Source Code
COPY . .

# Get Dependencies
RUN go mod download

# Build Binary
RUN go build -o ./out/simple-kv

# Give Binary Execution Permissions
RUN chmod +x ./out/simple-kv

# Wait for Build to Finish
RUN echo "Build Finished"

# List Files in Build Image
RUN ls -la ./out

# Check if Binary is Executable
RUN ls -la ./out/simple-kv



# Run Image is Alpine 3.14.0
FROM alpine:3.17

# Define Working Directory
WORKDIR /app

# Copy Binary with proper Permissions and Ownership
COPY --from=build /src/out/simple-kv .

# Run Binary
CMD ["./simple-kv"]


