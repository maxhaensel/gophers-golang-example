/*
  Add Mutation for todos here
*/

package todoresolver

import (
	"context"
	"graphql/model"
)

// CreateTodo Mutation
func (r *ResolverTodo) CreateTodo(ctx context.Context, args *struct {
	Todo string
}) (*ResponseTodo, error) {
	todo := &model.Todo{
		UID:  "id 2",
		Name: args.Todo,
	}
	return &ResponseTodo{todo: todo}, nil
}
