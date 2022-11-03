package server

import (
	"flag"
	"net"
	"sync"

	gRPC "github.com/bemillant/gRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	gRPC.UnimplementedChatServiceServer
	name  string
	port  string
	mutex sync.Mutex
}

var serverName = flag.String("name", "default", "Senders name")
var port = flag.String("port", "5400", "Server port")

func main() {
	listener, _ := net.Listen("tcp", "localhost:5400")

	var opts []grpc.ServerOption
	opts = append(opts,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	grpcServer := grpc.NewServer(opts)

	server := &Server{
		name: "Test Server 1",
		port: *port,
	}

	gRPC.RegisterChatServiceServer(grpcServer, server.UnimplementedChatServiceServer)
	grpcServer.Serve(listener)

	conn, err := grpc.Dial(":5400", opts)
	//stream, err := server.SendMessage(context.Background())
}

func (s *Server) BroadCastMessage(msgStream gRPC.ChatService_SendMessageServer) error {
	for {
		msg, err := msgStream.Recv()
		if err != nil {
			break
		}
	}
	ack := gRPC.MessageAck{
		Message: "Test",
	}
	msgStream.SendAndClose(&ack)
	return nil
}

func (s *Server) SendMessage(msg gRPC.ChatService_SendMessageClient) (gRPC.MessageAck, error) {
	ack := gRPC.MessageAck{
		Message: "Message was sent to server",
	}
	return ack, nil
}
