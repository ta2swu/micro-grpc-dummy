package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"context"

	proto "dummystream/bidirectionaldummy"

	"github.com/micro/go-micro"
)

func BidirectionalStream(cl proto.DummyStreamService) {
	// create streaming client
	ctx := context.Background()
	stream, err := cl.Stream(ctx)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	done := make(chan bool)
	// bidirectional stream
	// send  messages for a 20 count

	go func() {
		for j := 0; j < 30; j++ {

			fmt.Println(" send message : ", j)
			time.Sleep(1000 * time.Millisecond)

			message := "client message : " + strconv.Itoa(j)

			err := stream.Send(&proto.Request{Reqmsg: message})
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not send %v", err)
			}
		}
	}()
	go func() {
		for {

			rsp, err := stream.Recv()
			if err != nil {
				fmt.Println("recv err", err)
				break
			}
			fmt.Printf("\n got msg :", rsp.Resmsg)
		}
	}()
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done

	fmt.Println("streaming done")

}

func main() {
	//consulReg := cReg.NewRegistry(registry.Addrs("http://127.0.0.1:80"))

	service := micro.NewService(
	//micro.Name("dummy.stream.client"),
	//micro.Transport(trns),
	//micro.Registry(consulReg),
	)
	service.Init()

	// create client
	cl := proto.NewDummyStreamService("dummy.stream", service.Client())
	// bidirectional-stream
	time.Sleep(3 * time.Second)
	BidirectionalStream(cl)

}
