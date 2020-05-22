package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	"github.com/inklabs/rangedb/pkg/grpc/rangedbpb"
)

func main() {
	aggregateTypes := flag.String("aggregateTypes", "", "aggregateTypes separated by comma")
	host := flag.String("host", "127.0.0.1:8081", "RangeDB gRPC host address")
	flag.Parse()

	// TODO
	if *aggregateTypes != "" {
		log.Fatalf("not yet supported")
	}

	conn, err := grpc.Dial(*host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to dial (%s): %v", *host, err)
	}

	rangeDBClient := rangedbpb.NewRangeDBClient(conn)
	ctx := context.Background()
	eventsRequest := &rangedbpb.EventsRequest{
		StartingWithEventNumber: 0,
	}

	events, err := rangeDBClient.SubscribeToEvents(ctx, eventsRequest)
	if err != nil {
		log.Fatalf("unable to get events: %v", err)
	}

	for {
		record, err := events.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving event: %v", err)
		}

		fmt.Println(record)
	}
}