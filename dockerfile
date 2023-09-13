FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY proto ./proto

RUN apt-get update
RUN export PATH=$PATH:/usr/local/go/bin
RUN apt-get install -y protobuf-compiler
RUN go get google.golang.org/grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN export PATH="$PATH:$(go env GOPATH)/bin"

RUN protoc --go_out=./proto --go_opt=paths=import \ 
--go-grpc_out=./proto --go-grpc_opt=paths=import \
 ./proto/*.proto

RUN go build -o bin .

RUN go get github.com/rabbitmq/amqp091-go

# RUN protoc --go_out=./proto ./proto/*.proto

# RUN go run server_regional.go
# RUN go run server_central.go

ENTRYPOINT [ "/app/bin" ]