package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	pb "github.com/xxsxa/go-grpc-microservices-stock-service/proto/stock"
	"io/ioutil"
	"log"
	"os"
)

const (
	address         = "localhost:5051"
	defaultFilename = "stock.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Did not connet: %v", err)
	}
	json.Unmarshal(data, &consignment)
	return consignment, err

}

func main() {
	service := micro.NewService(micro.Name(""))
	service.Init()
	client := pb.NewShippingServiceClient("shippy.service.stock", service.Client())
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file : %v", err)
	}
	r, err := client.CreateConsignment(context.Background(), consignment)
	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
	log.Printf("Created: %t", r.Created)
}
