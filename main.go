package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	_ = AccountChange{
		Account: "0x7FEbaDA1daEFdA307c6F52f4818De6fE190C5B82",
	}
	client := conn.NewClie8nt()

	defer conn.Close()

	//client := pb.NewInventoryClient(conn)
	//bookList, err := client.GetBookList(context.Background(), &pb.GetBookListRequest{})
	//if err != nil {
	//	log.Fatalf("failed to get book list: %v", err)
	//}
	//log.Printf("book list: %v", bookList)
}
