# Start with a golang image
FROM golang:alpine
# Set a working directory
WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy .go files
COPY *.go ./

# Build 
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go

# Expose Application on port
EXPOSE 8090

# Run binary
CMD [ "./main" ]