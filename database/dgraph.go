package database

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

func NewClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

func Setup(c *dgo.Dgraph) error {
	// Install a schema into dgraph. Accounts have a `name` and a `balance`.
	err := c.Alter(context.Background(), &api.Operation{
		Schema: `
			id: int @index(term) .
			name: string .

			type Node {
				id
				name
				data
				class
				html
				typenode
			}
		`,
	})

	if err != nil {
		return err
	}
	return nil
}
