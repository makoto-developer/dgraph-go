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
	ctx := context.Background()
	dc, cancel := getDgraphClient()
	defer cancel()

	user := User{
		Uid:     "_:m",
		Name:    "yugi mutoh",
		Description:     "I am the king of duelist.",
		Tweets: []Tweet{
			{
				Content: "Duel!!!",
				Public: true,
				Like: []User{},
				DType: []string{"Tweet"},
			},
		},
		Follow: []User{},
		DType: []string{"User"},
	}
	addUser(ctx, dc, user)
	log.Println("finish!")
}

func addUser(ctx context.Context, dc *dgo.Dgraph, user User) {

	mutation := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	mutation.SetJson = pb
	response, err := dc.NewTxn().Mutate(ctx, mutation)
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

