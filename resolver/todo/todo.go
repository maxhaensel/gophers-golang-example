/*
  this file contains all Field-Resolvers
*/

package todoresolver

import (
	"context"
	"graphql/model"
	commentresolver "graphql/resolver/comment"

	"github.com/graph-gophers/dataloader"
	"github.com/graph-gophers/graphql-go"
)

// ResolverTodo export
type ResolverTodo struct{}

// ResponseTodo to resolve the Todo-Model
type ResponseTodo struct {
	todo *model.Todo
}

// UID resolver
func (r *ResponseTodo) UID() *graphql.ID {
	uid := graphql.ID(r.todo.UID)
	return &uid
}

// Name resolver
func (r *ResponseTodo) Name() *string {
	return &r.todo.Name
}

/*
Comment resolver
comments resolved via Dataloader
*/
func (r *ResponseTodo) Comment(ctx context.Context, args struct {
	ID *string
}) *[]*commentresolver.ResponseComment {
	thunk := commentresolver.CommentDataLoader.Load(ctx, dataloader.StringKey(r.todo.UID))
	result, err := thunk()
	if err != nil {
		// TODO: Fix err
		panic(err)
	}
	return result.(*[]*commentresolver.ResponseComment)
}
