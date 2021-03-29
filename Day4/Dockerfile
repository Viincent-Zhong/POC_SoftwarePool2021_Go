FROM golang:alpine
ENV FOO/=HELLO_MESSAGE
EXPOSE 8080
COPY exo1 .
WORKDIR ./exo1
RUN go mod download
RUN go run main.go