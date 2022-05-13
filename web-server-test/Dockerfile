FROM golang:1.18-alpine

WORKDIR /server

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN mkdir /server/static
COPY static/*.html ./static/
COPY *.go ./

RUN go build -o server

EXPOSE 8080
CMD [ "./server" ]
