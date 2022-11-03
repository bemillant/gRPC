package server

import (
	"flag"
	"net"
	"sync"

	gRPC "github.com/bemillant/gRPC"

	"google.golang.org/grpc"
)

type Server struct {
	gRPC.UnimplementedTemplateServer
	name  string
	port  string
	mutex sync.Mutex
}

var serverName = flag.String("name", "default", "Senders name")
var port = flag.String("port", "5400", "Server port")

func main() {
	listener, _ := net.Listen("tcp", "localhost:5400")

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts)

	server := &Server{
		name: "Test Server 1",
		port: *port,
	}

}

func (s *Server) SendMessage(msg gRPC.Message) (gRPC.MessageAck, error) {

	return nil, nil
}
