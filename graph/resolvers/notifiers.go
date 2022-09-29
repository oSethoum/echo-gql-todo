package resolvers

import (
	"todo/ent"
	"todo/graph/models"
)

func NotifyUsersListenners(r *mutationResolver, event models.Event) {
	r.UsersListennersMutext.Lock()
	for key := range r.UsersListenners {
		if r.UsersListenners[key].Event == event {
			println("Client: ", key, "\nEvent: ", event)
			query := r.UsersListenners[key].Query
			if query == nil {
				query = &models.UsersQueryInput{}
			}

			entities, err := r.Client.User.Query().Paginate(r.UsersListenners[key].Context, query.After, query.First, query.Before, query.Last, ent.WithUserFilter(query.Where.Filter), ent.WithUserOrder(query.OrderBy))

			if err == nil {
				r.UsersListenners[key].Channel <- entities
			}

		}
	}
	r.UsersListennersMutext.Unlock()
}

func NotifyUserListenners(r *mutationResolver, event models.Event, entity *ent.User) {
	r.UserListennersMutext.Lock()
	for key := range r.UserListenners {
		if r.UserListenners[key].Event == event && r.UserListenners[key].ID == entity.ID {
			r.UserListenners[key].Channel <- entity
		}
	}
	r.UserListennersMutext.Unlock()
}

func NotifyTodosListenners(r *mutationResolver, event models.Event) {
	r.TodosListennersMutext.Lock()
	for key := range r.TodosListenners {
		if r.TodosListenners[key].Event == event {
			println("Client: ", key, "\nEvent: ", event)
			query := r.TodosListenners[key].Query
			if query == nil {
				query = &models.TodosQueryInput{}
			}

			entities, err := r.Client.Todo.Query().Paginate(r.TodosListenners[key].Context, query.After, query.First, query.Before, query.Last, ent.WithTodoFilter(query.Where.Filter), ent.WithTodoOrder(query.OrderBy))

			if err == nil {
				r.TodosListenners[key].Channel <- entities
			}

		}
	}
	r.TodosListennersMutext.Unlock()
}

func NotifyTodoListenners(r *mutationResolver, event models.Event, entity *ent.Todo) {
	r.TodoListennersMutext.Lock()
	for key := range r.TodoListenners {
		if r.TodoListenners[key].Event == event && r.TodoListenners[key].ID == entity.ID {
			r.TodoListenners[key].Channel <- entity
		}
	}
	r.TodoListennersMutext.Unlock()
}
