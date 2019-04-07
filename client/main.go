package main

import (
	"fmt"
	"log"

	"github.com/HectorBarrios/grpcdemo-pluralsight-server/pb"
	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

const port = ":9000"

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewEmployeeServiceClient(conn)
	SendMetadata(client)
	fmt.Println("holi")
}

func SendMetadata(client pb.EmployeeServiceClient) {
	md := metadata.MD{}
	md["user"] = []string{"wakawaka"}
	md["password"] = []string{"12345678"}

	ctx := context.Background()
	ctx = metadata.NewIncomingContext(ctx, md)
	client.GetByBadgeNumber(ctx, &pb.GetByBadgeNumberRequest{})
}
