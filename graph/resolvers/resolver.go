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

type Resolver struct {
	Client *ent.Client

	UserListenners        map[string]UserListenner
	UserListennersMutext  sync.Mutex
	UsersListenners       map[string]UsersListenner
	UsersListennersMutext sync.Mutex

	TodoListenners        map[string]TodoListenner
	TodoListennersMutext  sync.Mutex
	TodosListenners       map[string]TodosListenner
	TodosListennersMutext sync.Mutex
}

var schema *graphql.ExecutableSchema

func ExecutableSchema() graphql.ExecutableSchema {
	if schema == nil {
		schema = new(graphql.ExecutableSchema)
		*schema = generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
			Client: db.DB,

			UserListenners:        make(map[string]UserListenner),
			UserListennersMutext:  sync.Mutex{},
			UsersListenners:       make(map[string]UsersListenner),
			UsersListennersMutext: sync.Mutex{},

			TodoListenners:        make(map[string]TodoListenner),
			TodoListennersMutext:  sync.Mutex{},
			TodosListenners:       make(map[string]TodosListenner),
			TodosListennersMutext: sync.Mutex{},
		}})
	}

	return *schema
}
