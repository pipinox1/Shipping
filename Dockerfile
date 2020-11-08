# Start from go latest image
FROM golang:latest

# Setting work directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependcy
RUN go mod download

# Copy code into work directory
COPY . .

# Building binary file
RUN go build -o main .

# Exposer 8080 port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
