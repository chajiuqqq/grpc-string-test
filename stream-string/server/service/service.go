package service

import (
	"context"
	"errors"
	"grpc-string-test/pb"
	. "grpc-string-test/pb"
	"io"
	"log"
	"strings"
)

const (
	StrMaxSize = 1024
)

// Service errors
var (
	ErrMaxSize = errors.New("maximum size of 1024 bytes exceeded")

	ErrStrValue = errors.New("maximum size of 1024 bytes exceeded")
)

type StringService struct {
	pb.UnimplementedStringServiceServer
}

func (s *StringService) Concat(ctx context.Context, req *StringRequest) (*StringResponse, error) {
	log.Println(req)
	if len(req.A)+len(req.B) > StrMaxSize {
		response := StringResponse{Ret: ""}
		return &response, nil
	}
	response := StringResponse{Ret: req.A + req.B}
	return &response, nil
}

func (s *StringService) Diff(ctx context.Context, req *StringRequest) (*StringResponse, error) {
	log.Println(req)
	if len(req.A) < 1 || len(req.B) < 1 {
		response := StringResponse{Ret: ""}
		return &response, nil
	}
	res := ""
	if len(req.A) >= len(req.B) {
		for _, char := range req.B {
			if strings.Contains(req.A, string(char)) {
				res = res + string(char)
			}
		}
	} else {
		for _, char := range req.A {
			if strings.Contains(req.B, string(char)) {
				res = res + string(char)
			}
		}
	}
	response := StringResponse{Ret: res}
	return &response, nil
}

//服务端返回流：接受一个req，一个qs
func (s *StringService) ConcatServerStream(req *StringRequest, qs StringService_ConcatServerStreamServer) error {
	response := StringResponse{Ret: req.A + req.B}
	for i := 0; i < 10; i++ {
		qs.Send(&response)
	}
	return nil

}

//客户端发流
func (s *StringService) ConcatClientStream(qs StringService_ConcatClientStreamServer) error {
	var params []string
	for {
		item, err := qs.Recv()
		//客户端发完了，服务器才处理并返回一个response
		if err == io.EOF {
			qs.SendAndClose(&StringResponse{Ret: strings.Join(params, "")})
			return nil
		}
		if err != nil {
			log.Println("fail to recv:", err)
			return err
		}
		params = append(params, item.A, item.B)
	}
}
func (s *StringService) ConcatDoubleStream(qs StringService_ConcatDoubleStreamServer) error {
	for {
		in, err := qs.Recv()
		//客户端发完了，服务器才处理并返回一个response
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Println("fail to recv:", err)
			return err
		}
		qs.Send(&StringResponse{Ret: in.A + in.B})
	}
}
