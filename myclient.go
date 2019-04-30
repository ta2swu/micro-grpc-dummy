package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"

	//proto "github.com/micro/examples/stream/server/proto"

	pb "servicestt/exproto"

	"google.golang.org/grpc"
)

type ourStreamer struct {
}

func main() {

	rand.Seed(time.Now().Unix())

	// dail server
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := pb.NewMystreamClient(conn)
	stream, err := client.Exstream(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	ctx := stream.Context()
	done := make(chan bool)

	// first goroutine sends random increasing numbers to stream
	// and closes int after 10 iterations
	go func() {
		for i := 1; i <= 100; i++ {
			// generate random nummber and send it to stream
			rnd := rand.Intn(i)

			str := strconv.Itoa(rnd)
			str = "this is message : " + str
			req := pb.Request{Reqmsg: str}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			log.Printf(" SENT MESSAGE : ", req.Reqmsg)
			time.Sleep(time.Millisecond * 200)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// second goroutine receives data from stream
	// and saves result in max variable
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			reponsemessage := resp.Resmsg
			log.Printf("received: ", reponsemessage)
		}
	}()

	// third goroutine closes done channel
	// if context is done
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
}
