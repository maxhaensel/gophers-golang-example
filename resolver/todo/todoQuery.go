/*
  Add Querys for todos here
*/

package todoresolver

import (
	"context"
	"graphql/model"
)

// Todo Resolver
func (r *ResolverTodo) Todo(ctx context.Context, args *struct {
	UID string
}) (*ResponseTodo, error) {
	todo := &ResponseTodo{todo: &model.Todo{Name: "name", UID: "id1"}}
	return todo, nil
}

// Todos Resolver
func (r *ResolverTodo) Todos(ctx context.Context, args *struct {
	UID string
}) (*[]*ResponseTodo, error) {
	arr := &[]*ResponseTodo{
		{todo: &model.Todo{Name: "Todo 1", UID: "id1"}},
		{todo: &model.Todo{Name: "Todo 2", UID: "id2"}},
		{todo: &model.Todo{Name: "Todo 3", UID: "id3"}},
	}
	return arr, nil
}
