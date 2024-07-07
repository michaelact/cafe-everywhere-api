FROM golang:1.22 as builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go build -o cafe-everywhere-api
ENTRYPOINT ["/app/cafe-everywhere-api"]
