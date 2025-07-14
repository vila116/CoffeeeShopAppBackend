// client interface with our service
package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/vila116/proto_example/coffee_Shop_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("falied to connect to grpc server :%s", err)
	}
	defer conn.Close()
	c := pb.NewCoffeeShopClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	menuStream, err := c.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatalf("error calling fucntion GetMenu() :%s", err)
	}
	done := make(chan bool)
	var items []*pb.Item
	go func() {
		for {
			resp, err := menuStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("error in recieving :%v", err)
			}
			items = resp.Items
			log.Printf("Resp received :%v", items)
		}

	}()
	<-done
	receipt, err := c.PlaceOrder(ctx, &pb.Order{Items: items})
	if err != nil {
		log.Fatalf("error while placing order%v", err)

	}
	log.Printf("%v", receipt)
	stauts, err := c.GetOrderStatus(ctx, receipt)
	if err != nil {
		log.Fatalf("error while getting order Status:%v", err)
	}
	log.Printf("%v", stauts)
}
