package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/adjoined-tb/ms.adjoined.me/src/asker/pbgo"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

type askerServer struct {
}

func main() {
	ctx := context.Background()
	var greeterSvcAddr string
	var greeterSvcConn *grpc.ClientConn
	mustMapEnv(&greeterSvcAddr, "GREETER_SERVICE_ADDR")
	mustConnGRPC(ctx, &greeterSvcConn, greeterSvcAddr)
	defer greeterSvcConn.Close()
	c := pb.NewGreeterClient(greeterSvcConn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3))
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}
