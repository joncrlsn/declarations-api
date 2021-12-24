# Compile and Build
FROM golang:1.17 AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /declarations-api

# Deploy
FROM alpine:3.15
WORKDIR /
COPY --from=build /declarations-api /declarations-api

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/declarations-api"]
