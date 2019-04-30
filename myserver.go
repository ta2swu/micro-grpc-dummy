package main

import (
	"fmt"
	"log"
	"net"

	//proto "github.com/micro/examples/stream/server/proto"

	pb "servicestt/exproto"

	"google.golang.org/grpc"
)

type ourStreamer struct {
}

func (s *ourStreamer) Exstream(streamServer pb.Mystream_ExstreamServer) error {

	stream := streamServer
	fmt.Println("should print only one time ")
	//ctx := stream.Context()
	intchan := make(chan string)
	go gomessage(stream, intchan)
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		mybyte := req.Reqmsg
		intchan <- mybyte
		log.Printf("Got msg %v", req.Reqmsg)

	}

	//fmt.Println("streming done")
	//
	return nil

}

func gomessage(stream pb.Mystream_ExstreamServer, intchan chan string) {

	for {

		m := <-intchan

		if err := stream.Send(&pb.Response{Resmsg: m}); err != nil {
			fmt.Println(" error sending", err)
		}
		fmt.Println("server sent message: ", m)
	}
}

// type MystreamServer interface {
// 	Exstream(Mystream_ExstreamServer) error
// }

func main() {

	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("server initilized ")
	// create grpc server
	s := grpc.NewServer()
	pb.RegisterMystreamServer(s, &ourStreamer{})

	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	/*if err := service.Run(); err != nil {
		log.Fatal(err)
	}*/
}
