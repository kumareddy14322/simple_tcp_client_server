package main
import (
	"log"
	 pb "github.com/kumareddy14322/simple_tcp_client_server/grpc2/datafiles/datafiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
address = "localhost:50051"
)

func main() {
// Set up a connection to the server.
conn, err := grpc.Dial(address, grpc.WithInsecure())
if err != nil {
	log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMoneyTransactionClient(conn)
	// Prepare data. Get this from clients like Frontend or App
	from := "0000"
	to := "5678"
	amount := float32(1250.75)
	// Contact the server and print out its response.
	r, err := c.MakeTransaction(context.Background(),
	&pb.TransactionRequest{From: from,
	To: to, Amount: amount})
	if err != nil {
	log.Fatalf("Could not transact: %v", err)
	}
	log.Printf("Transaction confirmed: %t", r.Confirmation)
}