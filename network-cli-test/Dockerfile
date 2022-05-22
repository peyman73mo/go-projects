FROM golang:1.17-alpine

WORKDIR /cli
RUN mkdir /cli/cmd
COPY cmd/cmd.go ./cmd
COPY ["cmd.go","go.mod","go.sum","./"]
RUN go mod download

RUN go build -o cli cmd.go
