PATH=$PATH:$GOPATH/bin
protodir=../../pb

protoc --go_out=./pbgo --go_opt=paths=source_relative \
    --go-grpc_out=./pbgo --go-grpc_opt=paths=source_relative \
    -I $protodir $protodir/adjoined.proto
