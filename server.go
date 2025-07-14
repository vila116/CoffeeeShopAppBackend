package main

import (
	"context"
	"log"
	"net"

	pb "github.com/vila116/proto_example/coffee_Shop_proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest, srv pb.CoffeeShop_GetMenuServer) error {
	Menuitems := []*pb.Item{
		&pb.Item{
			Id:   "1",
			Name: "Black Coffee",
		},
		&pb.Item{
			Id:   "2",
			Name: "Americano",
		},
		&pb.Item{
			Id:   "3",
			Name: "Vanilla Soy Chai Latte",
		},
	}
	for i, _ := range Menuitems {
		srv.Send(&pb.Menu{
			Items: Menuitems[0 : i+1],
		})
	}
	return nil
}
func (s *server) PlaceOrder(ctx context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "ABDC123",
	}, nil
}
func (s *server) GetOrderStatus(ctx context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "IN PROGRESS",
	}, nil
}

func main() {
	// setup a listener on port 9001
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("error while listening..:%v", err)
	}
	grpcServer := grpc.NewServer()
	//register server
	pb.RegisterCoffeeShopServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve:%s", err)
	}
}
