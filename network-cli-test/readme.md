terminal:

    print help:
        go run cmd.go help

    get request:
        go run cmd.go get --url [url]
        go run cmd.go get --url [url] --findIP true     // + DNS loopup

    DNS lookup:
        go run cmd.go dnslookup --url [url]

Docker:

    docker build --tag network-cli-test .   // build docker image
    docker run -it network-cli-test