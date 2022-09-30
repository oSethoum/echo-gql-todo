package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"todo/ent"
	"todo/graph/generated"
	"todo/graph/models"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
	entity, err := ent.FromContext(ctx).Todo.Create().SetInput(input).Save(ctx)

	if err == nil {
		go NotifyTodoListenners(r, models.EventCreate, entity)
		go NotifyTodosListenners(r, models.EventCreate)
	}

	return entity, err
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id int, input ent.UpdateTodoInput) (*ent.Todo, error) {
	entity, err := ent.FromContext(ctx).Todo.UpdateOneID(id).SetInput(input).Save(ctx)

	if err == nil {
		go NotifyTodoListenners(r, models.EventUpdate, entity)
		go NotifyTodosListenners(r, models.EventUpdate)
	}

	return entity, err
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id int) (*ent.Todo, error) {
	entity, err := r.Client.Todo.Get(ctx, id)

	if err == nil {
		go NotifyTodoListenners(r, models.EventDelete, entity)
		go NotifyTodosListenners(r, models.EventDelete)
	}

	return entity, ent.FromContext(ctx).Todo.DeleteOneID(id).Exec(ctx)
}

// Todo is the resolver for the Todo field.
func (r *queryResolver) Todo(ctx context.Context, id int) (*ent.Todo, error) {
	return r.Client.Todo.Get(ctx, id)
}

// Todos is the resolver for the todos field.
func (r *subscriptionResolver) Todos(ctx context.Context, event models.Event, query *models.TodosQueryInput) (<-chan *ent.TodoConnection, error) {
	clientId := ID()

	channel := make(chan *ent.TodoConnection, 1)
	println("----------------------------------------------------")
	println("Socket Client: ", clientId)
	println("Entity: Todos")
	println("Event: ", event)
	println("----------------------------------------------------")

	r.TodosListennersMutext.Lock()
	r.TodosListenners[clientId] = TodosListenner{
		Context: ctx,
		Channel: channel,
		Event:   event,
		Query:   query,
	}
	r.TodosListennersMutext.Unlock()

	go func() {
		<-ctx.Done()
		println("----------------------------------------------------")
		println("Socket Client: ", clientId)
		println("Disconnected")
		println("----------------------------------------------------")
		r.TodosListennersMutext.Lock()
		delete(r.TodosListenners, clientId)
		r.TodosListennersMutext.Unlock()
	}()

	return channel, nil
}

// Todo is the resolver for the Todo field.
func (r *subscriptionResolver) Todo(ctx context.Context, event models.Event, id int) (<-chan *ent.Todo, error) {
	clientId := ID()
	channel := make(chan *ent.Todo, 1)
	println("----------------------------------------------------")
	println("Socket Client: ", clientId)
	println("Entity: Todos")
	println("Event: ", event)
	println("Disconnected")
	println("----------------------------------------------------")

	r.TodoListennersMutext.Lock()
	r.TodoListenners[clientId] = TodoListenner{
		Context: ctx,
		ID:      id,
		Channel: channel,
		Event:   event,
	}
	r.TodoListennersMutext.Unlock()

	// remove listenner when the socket is disconnected
	go func() {
		<-ctx.Done()
		println("----------------------------------------------------")
		println("Socket Client: ", clientId)
		println("Disconnected")
		println("----------------------------------------------------")

		r.TodoListennersMutext.Lock()
		delete(r.TodoListenners, clientId)
		r.TodoListennersMutext.Unlock()
	}()

	return channel, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
