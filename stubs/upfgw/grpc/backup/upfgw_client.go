package grpc

import (
	context "context"
	"flag"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/grpc/protos"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	// name = flag.String("name", defaultName, "Name to greet")
)

func UpfGwClient() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// c := pb.NewGreeterClient(conn)
	c2 := protos.NewN1N2MessageClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c2.SendN1N2MessageTransferData(ctx, &protos.N1N2MessageTransferDataRequest{
		PduSessionId: 22,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Resp: %d", r.GetPduSessionId())
}
