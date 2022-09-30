package resolvers

import (
	"context"
	"todo/ent"
	"todo/graph/models"
)

type TodoListenner struct {
	Context context.Context
	ID      int
	Event   models.Event
	Channel chan *ent.Todo
}

type TodosListenner struct {
	Context context.Context
	Channel chan *ent.TodoConnection
	Event   models.Event
	Query   *models.TodosQueryInput
}
