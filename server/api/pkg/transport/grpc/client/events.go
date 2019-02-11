package client

import (
	"context"
	"log"
	"time"

	grpc "google.golang.org/grpc"

	pb_events "rocket-server/server/events/pkg/transport/grpc/pb"
)

type EventsClient struct {
	conn    *grpc.ClientConn
	service pb_events.EventsClient
}

func NewEventsClient(url string) (*EventsClient, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb_events.NewEventsClient(conn)
	return &EventsClient{conn, c}, nil
}

func (c *EventsClient) Close() {
	c.conn.Close()
}
