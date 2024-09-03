FROM golang:1.22-buster

RUN go version
ENV GOPATH=/

COPY ././

RUN go mod download
RUN go build -o blog-app ./cmd/main.go
CMD ["./blog-app"]