package main

import (
    "context"
    "log"
    "net"

    greeter "gtihub.com/wso2/choreo-sample-apps/go/grpc-greeter/pkg"
    "google.golang.org/grpc"
)

type server struct{
    greeter.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloReply, error) {
    log.Printf("Received: %v", in.GetName())
    return &greeter.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    greeter.RegisterGreeterServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
