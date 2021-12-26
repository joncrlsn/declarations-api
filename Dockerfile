# Compile and Build
FROM golang:1.17 AS build
RUN mkdir /app
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /declarations-api

##########
# Deploy #
##########
FROM debian:latest
#RUN apk add curl \
# && adduser -S -u 1122 app \
RUN useradd --uid 1122 app \
  && mkdir /app \
  && mkdir /app/declarations \
  && mkdir /app/static 
COPY --from=build /declarations-api /app/declarations-api
COPY ./static/* /app/static/
RUN chown -R app /app 
USER app
EXPOSE 8080
WORKDIR /app

#ENV DECLARATIONS_FILE=/app/static/declarations
#ENV STATIC_DIR=/app/static
ENTRYPOINT exec /app/declarations-api
