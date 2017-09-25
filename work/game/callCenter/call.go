package callCenter

import (
	klog "goklmmx/lib/log"
	kconf "goklmmx/lib/conf"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "goklmmx/lib/pb3"
)

var callClient pb.GreeterClient

func CallCenterInit()  {
	kconf.SetFile("conf/config.cfg")
	klog.SetLogfile()
	c, _ := kconf.GetConf()
	address,_ := c.String("centerserver","addr")
	klog.Klog.Printf("centerserver addr :%s", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		klog.Klog.Printf("did not connect: %v", err)
	}
	//defer conn.Close()
	callClient = pb.NewGreeterClient(conn)
}

func Test()  {
	name := "heyuanlong"

	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()
	r, err := callClient.SayTest(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		klog.Klog.Printf("could not greet: %v", err)
	}else {
		klog.Klog.Printf("Greeting: %s", r.Message)
	}
}