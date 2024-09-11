package main

import (
	"context"
	"log"
	"net"
)

// 服务接口
type YourServiceInterface interface {
	YourMethod(ctx context.Context, req YourRequest) (YourResponse, error)
}

// 请求消息
type YourRequest struct {
	Field1 string
	Field2 string
}

// 响应消息
type YourResponse struct {
	Field1 string
	Field2 string
}

// 实现服务接口
type yourService struct{}

func (s *yourService) YourMethod(ctx context.Context, req YourRequest) (YourResponse, error) {
	// 这里实现你的业务逻辑
	res := YourResponse{
		Field1: "responseField1",
		Field2: "responseField2",
	}
	return res, nil
}

// 创建服务实例
func NewYourService() YourServiceInterface {
	return &yourService{}
}

func main() {
	// 创建服务实例
	s := NewYourService()

	// 创建监听器
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	// 接受连接
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Failed to accept: %v", err)
	}
	defer conn.Close()

	// 创建 gRPC 客户端
	client := YourServiceInterface(s)

	// 读取请求
	req := YourRequest{}
	err = conn.Read(&req)
	if err != nil {
		log.Fatalf("Failed to read request: %v", err)
	}

	// 调用方法
	res, err := client.YourMethod(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call method: %v", err)
	}

	// 发送响应
	err = conn.Write(res)
	if err != nil {
		log.Fatalf("Failed to write response: %v", err)
	}
}
