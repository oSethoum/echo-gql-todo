package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"todo/auth"
	"todo/ent"
	"todo/graph/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	entity, err := ent.FromContext(ctx).User.Create().SetInput(input).Save(ctx)

	if err == nil {
		go NotifyUserListenners(r, models.EventCreate, entity)
		go NotifyUsersListenners(r, models.EventCreate)
	}

	return entity, err
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	entity, err := ent.FromContext(ctx).User.UpdateOneID(id).SetInput(input).Save(ctx)

	if err == nil {
		go NotifyUserListenners(r, models.EventUpdate, entity)
		go NotifyUsersListenners(r, models.EventUpdate)
	}

	return entity, err
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*ent.User, error) {
	entity, err := r.Client.User.Get(ctx, id)

	if err == nil {
		go NotifyUserListenners(r, models.EventDelete, entity)
		go NotifyUsersListenners(r, models.EventDelete)
	}

	return entity, ent.FromContext(ctx).User.DeleteOneID(id).Exec(ctx)
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, id int) (*ent.User, error) {
	return r.Client.User.Get(ctx, id)
}

// Users is the resolver for the users field.
func (r *subscriptionResolver) Users(ctx context.Context, event models.Event, query *models.UsersQueryInput) (<-chan *ent.UserConnection, error) {
	socketClient := ctx.Value(auth.ContextKey{Key: "socketClient"}).(string)
	channel := make(chan *ent.UserConnection, 1)
	println("----------------------------------------------------")
	println("Socket Client: ", socketClient)
	println("Entity: Users")
	println("Event: ", event)
	println("----------------------------------------------------")

	r.UsersListennersMutext.Lock()
	r.UsersListenners[socketClient] = UsersListenner{
		Context: ctx,
		Channel: channel,
		Event:   event,
		Query:   query,
	}
	r.UsersListennersMutext.Unlock()

	go func() {
		<-ctx.Done()
		println("----------------------------------------------------")
		println("Socket Client: ", socketClient)
		println("Disconnected")
		println("----------------------------------------------------")
		r.UsersListennersMutext.Lock()
		delete(r.UsersListenners, socketClient)
		r.UsersListennersMutext.Unlock()
	}()

	return channel, nil
}

// User is the resolver for the User field.
func (r *subscriptionResolver) User(ctx context.Context, event models.Event, id int) (<-chan *ent.User, error) {
	socketClient := ctx.Value(auth.ContextKey{Key: "socketClient"}).(string)
	channel := make(chan *ent.User, 1)
	println("----------------------------------------------------")
	println("Socket Client: ", socketClient)
	println("Entity: Users")
	println("Event: ", event)
	println("Disconnected")
	println("----------------------------------------------------")

	r.UserListennersMutext.Lock()
	r.UserListenners[socketClient] = UserListenner{
		Context: ctx,
		ID:      id,
		Channel: channel,
		Event:   event,
	}
	r.UserListennersMutext.Unlock()

	// remove listenner when the socket is disconnected
	go func() {
		<-ctx.Done()
		println("----------------------------------------------------")
		println("Socket Client: ", socketClient)
		println("Disconnected")
		println("----------------------------------------------------")

		r.UserListennersMutext.Lock()
		delete(r.UserListenners, socketClient)
		r.UserListennersMutext.Unlock()
	}()

	return channel, nil
}
