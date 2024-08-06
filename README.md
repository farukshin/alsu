# alsu

## build server

```
protoc -I proto/ proto/alsu.proto --go_out=plugins=grpc:src/proto
cd src && go build -v -o ../build/server && cd ..

```

## build test client

```
protoc -I proto/ proto/alsu.proto --go_out=plugins=grpc:client/proto
cd client && go build -v -o ../build/client && cd ..

```