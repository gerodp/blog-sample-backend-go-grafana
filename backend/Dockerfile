FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the source code and dependencies
COPY . .

# Install dependencies
RUN go get -d -v ./...

# Build the application
RUN go build -o server .

# Expose the port on which the server listens
EXPOSE 9494

# Start the server when the container is run
ENTRYPOINT ["./server"]