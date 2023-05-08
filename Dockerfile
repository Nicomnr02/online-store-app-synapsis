# i make build stage first
FROM golang:1.19-alpine AS builder

RUN apk update && apk add --no-cache git
RUN rm -rf /app/*

WORKDIR /app
COPY . .
RUN go build -o app-store-server-synapsis main.go


# then running the stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/app-store-server-synapsis .

EXPOSE 8000
CMD ["/app/app-store-server-synapsis"]







