FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY proto ./proto
COPY parametros_de_inicio.txt .

RUN apt-get update
RUN export PATH=$PATH:/usr/local/go/bin
RUN apt-get install -y protobuf-compiler
#RUN go get github.com/streadway/amqp
RUN go get google.golang.org/grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN export PATH="$PATH:$(go env GOPATH)/bin"

#RUN protoc --go_out=. --go_opt=paths=source_relative \
#--go-grpc_out=. --go-grpc_opt=paths=source_relative \
# ./proto/*.proto
RUN go build -o bin .

# RUN protoc --go_out=./proto ./proto/*.proto

# RUN go run server_regional.go
# RUN go run server_central.go

ENTRYPOINT [ "/app/bin" ]