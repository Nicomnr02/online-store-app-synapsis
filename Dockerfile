FROM golang:1.19-alpine as development

RUN apk update && apk add --no-cache git

WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Install Reflex for development
RUN go install github.com/cespare/reflex@latest
# Expose port
EXPOSE 7000
# Start app
CMD reflex -g '*.go' go run main.go --start-service

