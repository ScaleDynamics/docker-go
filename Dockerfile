FROM golang:latest
WORKDIR /usr/src/app
COPY . .
RUN go build -o /usr/bin/app index.go
EXPOSE 8080
CMD ["/usr/bin/app"]