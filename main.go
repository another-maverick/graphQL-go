package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/friendsofgo/graphiql"
	gql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)


//defining schema here
const Schema = `
type Player { 
name: String!
team: String!
height: Int!
title: String!
 }

type Query {
 player(name: String!): Player
}

schema {
query: Query
}
`

//defining data model here
type Player struct {
	name string
	team string
	height int
	title string
}

func strToPtr(s string) *string {
	return &s
}

var players map[string]Player

type query struct{}

//Resolver here
type PlayerResolver struct {
	player *Player
}

func (pr *PlayerResolver) Name() string {
	return pr.player.name
}

func (pr *PlayerResolver) Team() string {
	return pr.player.team
}

func (pr *PlayerResolver) Height() int32 {
	return int32(pr.player.height)
}

func (pr *PlayerResolver) Title() string {
	return pr.player.title
}

// Query Resolver here. This returns the resolver
func (q *query) Player(ctx context.Context, args struct {Name string}) *PlayerResolver {
	pl, ok := players[strings.ToLower(args.Name)]
	if ok {
		return &PlayerResolver{player: &pl}
	}
	return nil
}

func main() {
	fmt.Println("Hello")

	schema := gql.MustParseSchema(Schema, &query{})
	http.Handle("/graphql", &relay.Handler{Schema: schema})

	// Initialize players here
	players = map[string]Player{
		"curry": {
			name: "Stephen Curry",
			team: "Warriors",
			height: 6,
			title: "Best Shooter",
		},
		"klay": {
			name: "Klay Thompson",
			team: "Warriors",
			height: 6,
			title: "Best Guard",
		},
		"lebron": {
			name: "LeBron James",
			team: "Lakers",
			height: 6,
			title: "Over Rated",
		},
	}

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}
	http.Handle("/", graphiqlHandler)


	log.Println("Strating the Server")
	log.Fatal(http.ListenAndServe(":12345", nil))


}
