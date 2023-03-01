package main

import (
	"context"
	"fmt"
	"grpc-string-test/pb"
	"io"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

func simpleConn(stringClient pb.StringServiceClient) {

	stringReq := &pb.StringRequest{A: "A", B: "B"}
	reply, _ := stringClient.Concat(context.Background(), stringReq)
	fmt.Printf("StringService Concat : %s concat %s = %s\n", stringReq.A, stringReq.B, reply.Ret)
}

//服务端返回流，客户端需要接受流
func streamServerConn(stringClient pb.StringServiceClient) {

	stringReq := &pb.StringRequest{A: "A", B: "B"}
	//接受流
	stream, _ := stringClient.ConcatServerStream(context.Background(), stringReq)
	//循环读取流
	for {
		item, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("fail to recv:", err)
		}
		fmt.Printf("StringService Concat : %s concat %s = %s\n", stringReq.A, stringReq.B, item.Ret)
	}

}

//客户端发送流
func streamClientConn(stringClient pb.StringServiceClient) {

	stream, err := stringClient.ConcatClientStream(context.Background())
	//流式发送requset
	for i := 0; i < 10; i++ {
		if err != nil {
			log.Println("fail to call:", err)
			break
		}
		stream.Send(&pb.StringRequest{A: strconv.Itoa(i), B: strconv.Itoa(i + 1)})
	}

	//发完后读取一个response
	recv, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("fail to recv:", err)
	}

	fmt.Printf("Ret is %s\n", recv.Ret)
}

//双向流
func doubleStreamConn(stringClient pb.StringServiceClient) {

	stream, _ := stringClient.ConcatDoubleStream(context.Background())
	var i int
	for {
		err := stream.Send(&pb.StringRequest{A: strconv.Itoa(i), B: strconv.Itoa(i + 1)})
		if err != nil {
			log.Println("fail to send:", err)
			break
		}
		recv, err := stream.Recv()
		if err != nil {
			log.Println("fail to recv:", err)
			break
		}
		fmt.Printf("Ret is %s\n", recv.Ret)
		i++
		time.Sleep(time.Second)
	}

}

func main() {
	serviceAddress := "127.0.0.1:1234"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()

	stringClient := pb.NewStringServiceClient(conn)
	// streamServerConn(stringClient)
	// streamClientConn(stringClient)
	doubleStreamConn(stringClient)

}
