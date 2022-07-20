package main


import(
	"context"
	"google.golang.org/grpc"
	"log"
	"mock-website/proto"
)

func main() {
	var conn *grpc.ClientConn

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

	log.Printf("Response from Server: %s", response.Body)

}