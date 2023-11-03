package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dgraph-io/dgo/v230"
	"github.com/dgraph-io/dgo/v230/protos/api"
	"google.golang.org/grpc"
)

type CancelFunc func()

type Tweet struct {
	Uid     string     	`json:"uid,omitempty"`
	Content	string     	`json:"Tweet.content,omitempty"`
	User 	User		`json:"Tweet.user,omitempty"`
	Public  bool		`json:"Tweet.public,omitempty"`
	Like    []User		`json:"Tweet.like,omitempty"`
	DType   []string	`json:"dgraph.type,omitempty"`
}

type User struct {
	Uid			string     	`json:"User.uid,omitempty"`
	Name    	string     	`json:"User.name,omitempty"`
	Description string     	`json:"User.description,omitempty"`
	Tweets		[]Tweet 	`json:"User.tweets,omitempty"`
	Follow		[]User		`json:"User.follow,iomitempty"`
	DType   	[]string   	`json:"dgraph.type,omitempty"`
}

func main() {
	dc, cancel := getDgraphClient()
	defer cancel()

	ctx := context.Background()
	u := User{
		Uid:     "_:m",
		Name:    "yuugi mutoh",
		Description:     "I am the king of duelist.",
		Tweets: []Tweet{{
			Content: "Duel!!!",
			Public: true,
			Like: []User{},
			DType:   []string{"Tweet"},
		}},
		Follow: []User{},
		DType:   []string{"User"},
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	response, err := dc.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
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
