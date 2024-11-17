package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"github.com/quic-go/quic-go/http3"

	greetv1 "github.com/tatsuya0429/http3_grpc_server/gen/greet/v1"
	greetv1connect "github.com/tatsuya0429/http3_grpc_server/gen/greet/v1/greetv1connect"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)
	addr := "127.0.0.1:8080"
	h3srv := http3.Server{
		Addr:    addr,
		Handler: mux,
	}
	if err := h3srv.ListenAndServeTLS("./secrets/server.crt", "./secrets/private.key"); err != nil {
		log.Fatal(err)
	}
}
