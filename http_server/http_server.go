package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"time"
)
import pb "github.com/alfssobsd/kube-grpc-example/baseproto"

func main() {
	//Get grpc server address from env `GRPC_SERVER_ADDR`
	address := os.Getenv("GRPC_SERVER_ADDR")

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("HTTP Server: %v\n", err)
	}
	defer conn.Close()
	myHostName, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := pb.NewGeneralServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		answer, err := c.GetServerName(ctx, &pb.ServerNameRequest{ClientName: myHostName})
		defer cancel()
		if err != nil {
			log.Fatalf("HTTP Server: could not get answer: %v\n", err)
		}
		log.Printf("HTTP Server: Request from %s to location = /, result = %s\n", r.RemoteAddr, answer.Name)
		fmt.Fprintf(w, "ServerName = %s", answer.Name)
	})
	log.Printf("HTTP Server: Start server on :8000 hostname = %s\n", myHostName)
	log.Printf("HTTP Server: GRPC connection to " + address + "\n")
	http.ListenAndServe(":8000", nil)
}
