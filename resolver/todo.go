/*
  this file contains all Field-Resolvers
*/

package resolver

import (
	"context"
	"graphql/model"

	"github.com/graph-gophers/dataloader"
	"github.com/graph-gophers/graphql-go"
)

// TodoResolver to resolve the Todo-Type
type TodoResolver struct {
	todo *model.Todo
}

// UID resolver
func (r *TodoResolver) UID() *graphql.ID {
	uid := graphql.ID(r.todo.UID)
	return &uid
}

// Name resolver
func (r *TodoResolver) Name() *string {
	return &r.todo.Name
}

/*
Comment resolver
comments resolved as Dataloader
*/
func (r *TodoResolver) Comment(ctx context.Context) *[]*CommentResolver {
	thunk := CommentDataLoader.Load(ctx, dataloader.StringKey(r.todo.UID))
	result, err := thunk()
	if err != nil {
		// TODO: Fix err
		panic(err)
	}
	return result.(*[]*CommentResolver)
}
