package main

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

type CancelFunc func()

func main() {
	dc, cancel := getDgraphClient()
	defer cancel()

	ctx := context.Background()
	op := &api.Operation{Schema: `
		Tweet.content: string .
	Tweet.content: string @index(fulltext) .
	Tweet.title: string .
	Tweet.user: uid .
	User.age: int @index(int) .
	User.name: string @index(hash) @upsert .
	User.tweets: [uid] .
	type Tweet {
		Tweet.title
		Tweet.content
		Tweet.user
	}
	type User {
		User.name
		User.age
		User.tweets
	}`,}
	err := dc.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}
}

func getDgraphClient() (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	dc := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	return dc, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}

