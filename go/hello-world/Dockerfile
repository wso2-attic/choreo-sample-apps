# Use an official Go runtime as a parent image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o main .

# Define the command to run your Go application
CMD ["./main"]
