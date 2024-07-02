# GRPC Installation

### Validate Go Installation
```bash
go version
```

### Export GOPATH
```bash
export GOPATH=$HOME/Work/go
echo $GOPATH
```

### Install Protoc
https://grpc.io/docs/protoc-installation/

### Install Protoc (for Windows)
https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/


### Validate Protoc Installation
```bash
protoc --version
```

### Install Protoc Dependecy for Golang
```bash
# protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
# proto
go get -u google.golang.org/protobuf/proto
# protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
# protoc-gen-swagger
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

### Validate Protoc Dependency Golang Installation
```bash
cd $GOPATH/bin
ls
```

### Create Golang Project and Create Proto Directory
```bash
├── project
│   ├── proto
```

### Download Necessary Library for Protobuf
```bash
├── project
│   ├── proto
│   │   ├── get-libs.sh
```
```bash
./get-libs.sh
```

### Validate Necessery Library was Download
```bash
cd proto/lib
ls
```

### Create First Proto File in proto/v1/apiname/apiname.proto
```bash
├── project
│   ├── proto
│   │   ├── get-libs.sh
│   │   ├── v1
│   │   │   ├── health
│   │   │   │   ├── health.proto
```

### Export GOPATH/bin to the PATH
```bash
export PATH=$PATH:$GOPATH/bin
```

### Compile Proto fil
```bash
cd proto
protoc \
      -I. \
      -I/usr/local/include \
      -I${GOPATH}/src \
      -I${GOPATH}/src/gitlab.com/go-grpc-learning/proto \
      -I${GOPATH}/src/gitlab.com/go-grpc-learning/proto/lib \
      --go_out=plugins=grpc:$GOPATH/src \
      --grpc-gateway_out=logtostderr=true:$GOPATH/src \
      v1/health/health.proto
```

## Validate Proto file Compilation
```bash
├── project
│   ├── proto
│   │   ├── get-libs.sh
│   │   ├── v1
│   │   │   ├── health
│   │   │   │   ├── health.proto
│   │   │   │   ├── health.pb.go
```

### Create main.go
```bash
├── project
│   ├── proto
│   ├── main.go
```

### go mod init & go mod tidy
```bash
go mod init
go mod tidy
```

### Validate go.mod & go.sum
```bash
├── project
│   ├── proto
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
```

### Create config.yaml
```bash
├── project
│   ├── proto
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── config.yaml
```

### Create configs module
```bash
├── project
│   ├── proto
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── config.yaml
│   ├── configs
│   │   ├── configs.go
```

### Create router/grpc.go
```bash
├── project
│   ├── proto
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── config.yaml
│   ├── configs
│   │   ├── configs.go
│   ├── router
│   │   ├── router.go
│   │   ├── grpc.go
```

### Add NewGRPCServer to main.go
```bash
g.Go(func() error { return router.NewGRPCServer(configs, logger) })
```

### Create api file go
```bash
├── project
│   ├── proto
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── config.yaml
│   ├── configs
│   ├── router
│   ├── api
│   │   ├── v1
│   │   │   ├── health
│   │   │   │   ├── health.go
│   │   │   │   ├── get.go
```

### Register api router/grpc.go
```bash
hlpb.RegisterHealthServiceServer(grpcServer, health.New(configs, logger))
```

### Run application
```bash
go run .
```

