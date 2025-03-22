FROM golang:1.24

WORKDIR /app
COPY . .

RUN go build -o command-server

EXPOSE 8080
CMD ["./command-server"]

