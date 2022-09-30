package resolvers

import (
	"todo/db"
	"todo/ent"
	"todo/graph/generated"

	"github.com/99designs/gqlgen/graphql"

	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var id = 0

func ID() int {
	id++
	return id
}

type Resolver struct {
	Client *ent.Client

	TodoListenners        map[int]TodoListenner
	TodoListennersMutext  sync.Mutex
	TodosListenners       map[int]TodosListenner
	TodosListennersMutext sync.Mutex
}

var schema *graphql.ExecutableSchema

func ExecutableSchema() graphql.ExecutableSchema {
	if schema == nil {
		schema = new(graphql.ExecutableSchema)
		*schema = generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
			Client: db.DB,

			TodoListenners:        make(map[int]TodoListenner),
			TodoListennersMutext:  sync.Mutex{},
			TodosListenners:       make(map[int]TodosListenner),
			TodosListennersMutext: sync.Mutex{},
		}})
	}

	return *schema
}
