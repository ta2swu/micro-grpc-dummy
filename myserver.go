package main

import (
	"fmt"
	"log"

	"context"

	proto "dummystream/bidirectionaldummy"

	"github.com/micro/go-micro"
)

type Streamer struct{}

func (e *Streamer) Stream(ctx context.Context, stream proto.DummyStream_StreamStream) error {

	mystream := stream
	fmt.Println("should print only one time ")
	//ctx := stream.Context()
	stringchan := make(chan string)

	i := 0
	go Sendmessage(stringchan, mystream)
	//(intchan, stream)
	for {

		i++
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		rmsg, err := mystream.Recv()
		if err != nil {
			return err
		}

		myessage := rmsg.Reqmsg
		stringchan <- myessage
		log.Println("message recieved ", myessage)
	}

	return nil
}

func Sendmessage(stringchan chan string, mystream proto.DummyStream_StreamStream) {

	for {

		message := <-stringchan
		message = "reply from server : " + message
		sendmsg := &proto.Response{Resmsg: message}
		if err := mystream.Send(sendmsg); err != nil {
			fmt.Println(" error sending", err)
		}
		fmt.Println("server sent message: ", message)

	}
}

func main() {
	// new service
	//consulReg := cReg.NewRegistry(registry.Addrs("http://127.0.0.1:80"))

	service := micro.NewService(
		micro.Name("dummy.stream"),
	// micro.Registry(consulReg),
	)

	// Init command line

	// Register Handler
	proto.RegisterDummyStreamHandler(service.Server(), new(Streamer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
