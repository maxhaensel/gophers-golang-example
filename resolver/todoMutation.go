/*
  Add Mutation for todos here
*/

package resolver

import (
	"context"
	"graphql/model"
)

// CreateTodo Mutation
func (r *Resolver) CreateTodo(ctx context.Context, args *struct {
	Todo string
}) (*TodoResolver, error) {
	todo := &model.Todo{
		UID:  "id 2",
		Name: args.Todo,
	}
	return &TodoResolver{todo: todo}, nil
}
