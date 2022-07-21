package main

import (
	"net/http"
	"google.golang.org/grpc"
	"mock-website/proto"
	"log"
)

func indexHandler(w http.ResponseWriter, r *http.Request, conn *grpc.ClientConn ) {
	conn,err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil{
		log.Fatalf("Error when calling Hello: %s", err)
	}
	defer conn.Close()

	c:= proto.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &proto.Message{Body: "Hello From Client!"})

	if err!= nil {
		log.Fatalf("Error when calling SayHello function rpc %s",err)
	}

	w.Write(byte[] "HELP ME")
}
