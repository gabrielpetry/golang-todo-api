FROM golang:alpine

WORKDIR /app

# Copy main app
COPY ./src /app/src
# Copy deps files
COPY ./go.mod ./go.sum ./modd.conf /app/

# reload source code
RUN go get github.com/cortesi/modd/cmd/modd

# install code deps
RUN go mod download

# ENTRYPOINT go run server/main.go
ENTRYPOINT modd
