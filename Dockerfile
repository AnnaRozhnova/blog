FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o blog ./cmd/main.go
CMD ["./blog"]