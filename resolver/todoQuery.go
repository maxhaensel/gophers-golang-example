/*
  Add Querys for todos here
*/

package resolver

import (
	"context"
	"graphql/model"
)

// Todo Resolver
func (r *Resolver) Todo(ctx context.Context, args *struct {
	UID string
}) (*TodoResolver, error) {
	todo := &TodoResolver{todo: &model.Todo{Name: "name", UID: "id1"}}
	return todo, nil
}

// Todos Resolver
func (r *Resolver) Todos(ctx context.Context, args *struct {
	UID string
}) (*[]*TodoResolver, error) {
	arr := &[]*TodoResolver{
		{todo: &model.Todo{Name: "Todo 1", UID: "id1"}},
		{todo: &model.Todo{Name: "Todo 2", UID: "id2"}},
		{todo: &model.Todo{Name: "Todo 3", UID: "id3"}},
	}
	return arr, nil
}
