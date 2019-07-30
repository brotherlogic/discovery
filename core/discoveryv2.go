package discovery

import (
	"context"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/discovery/proto"
	pbg "github.com/brotherlogic/goserver/proto"
)

type elector interface {
	elect(ctx context.Context, entry *pb.RegistryEntry) error
}

type prodElector struct {
}

func (p *prodElector) elect(ctx context.Context, entry *pb.RegistryEntry) error {
	conn, err := grpc.Dial(entry.Ip+":"+strconv.Itoa(int(entry.Port)), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	server := pbg.NewGoserverServiceClient(conn)
	_, err = server.Mote(ctx, &pbg.MoteRequest{Master: true})

	return err
}
