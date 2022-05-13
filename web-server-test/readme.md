
docker build --tag go-server .

docker run -p 8000:8080 --name go-test-server go-server