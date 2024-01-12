# Base image is alpine with Go installed
FROM golang:alpine

# Set the working directory to the app directory
WORKDIR /app

# Copy all files from the current directory to the app directory within the container
COPY . .

# Install dependencies and build the application
RUN go get && go build -o bin .

# Set the entrypoint to the built application
ENTRYPOINT ["/app/bin"]
