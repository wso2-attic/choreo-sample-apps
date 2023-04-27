package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	greeter "gtihub.com/wso2/choreo-sample-apps/go/grpc-greeter/pkg"
)

func main() {
	target := os.Getenv("GREETER_SERVICE")
	conn, err := grpc.Dial(target, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Did not connect: %v", err)
    }
    defer conn.Close()
    c := greeter.NewGreeterClient(conn)
    name := "Choreo"
    r, err := c.SayHello(context.Background(), &greeter.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.GetMessage())
}
