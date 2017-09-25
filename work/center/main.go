package main

import (
	"net"
	klog "goklmmx/lib/log"
	kconf "goklmmx/lib/conf"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "goklmmx/lib/pb3"
	"google.golang.org/grpc/reflection"
	//"time"
)
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
func (s *server) SayTest(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}


func main() {
	kconf.SetFile("conf/config.cfg")
	klog.SetLogfile()


	c, _ := kconf.GetConf()
	serverPort,_ := c.String("server","port")
	serverPort = ":"+serverPort
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		klog.Klog.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	klog.Klog.Printf("start")
	if err := s.Serve(lis); err != nil {
		klog.Klog.Printf("failed to serve: %v", err)
	}

}