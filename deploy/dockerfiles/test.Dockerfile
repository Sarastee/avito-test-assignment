FROM golang:1.22.1-alpine
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
ADD . .