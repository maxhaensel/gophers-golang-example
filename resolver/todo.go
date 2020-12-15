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

type todoArgs struct {
	UID string `json:"uid"`
}

// Todo Resolver
func (r *Resolver) Todo(ctx context.Context, args todoArgs) (*TodoResolver, error) {
	todo := &TodoResolver{todo: &model.Todo{Name: "name", UID: "id1"}}
	return todo, nil
}

// Todos Resolver
func (r *Resolver) Todos(ctx context.Context, args todoArgs) (*[]*TodoResolver, error) {
	arr := &[]*TodoResolver{
		{todo: &model.Todo{Name: "Todo 1", UID: "id1"}},
		{todo: &model.Todo{Name: "Todo 2", UID: "id2"}},
		{todo: &model.Todo{Name: "Todo 3", UID: "id3"}},
	}
	return arr, nil
}

// CreateTodo Mutation
func (r *Resolver) CreateTodo(ctx context.Context, args *struct {
	Todo string `json:"todo"`
}) (*TodoResolver, error) {
	todo := &model.Todo{
		UID:  "id 2",
		Name: args.Todo,
	}
	return &TodoResolver{todo: todo}, nil
}

// UID EXPORT
func (r *TodoResolver) UID() *graphql.ID {
	uid := graphql.ID(r.todo.UID)
	return &uid
}

// Name EXPORT
func (r *TodoResolver) Name() *string {
	return &r.todo.Name
}

// Comment Export
func (r *TodoResolver) Comment(ctx context.Context) *[]*CommentResolver {
	thunk := CommentDataLoader.Load(ctx, dataloader.StringKey(r.todo.UID))
	result, err := thunk()
	if err != nil {
		// TODO: Fix err
		panic(err)
	}

	return result.(*[]*CommentResolver)
}
