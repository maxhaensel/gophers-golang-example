package resolver

import (
	"context"
	"fmt"
	"graphql/model"

	"github.com/graph-gophers/dataloader"
)

// CommentResolver to resolve the Comment-Type
type CommentResolver struct {
	comment *model.Comment
}

// UID CommentResolver
func (r *CommentResolver) UID(ctx context.Context) *string {
	return &r.comment.UID
}

// Name CommentResolver
func (r *CommentResolver) Name(ctx context.Context) *string {
	return &r.comment.Name
}

// CommentDataLoader export a initialized instance to use in the Comment-Resolver
var CommentDataLoader = dataloader.NewBatchedLoader(batchFn)

/*
  Create a batchFunction to handle the n+1 Graphql-Problem
  keys are the properties they omited by the Comments-Resolver
*/
var batchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// Create a emptry result Array
	var results []*dataloader.Result
	for in, key := range keys {
		comment := &CommentResolver{comment: &model.Comment{
			UID:  key.String(),
			Name: "index: " + fmt.Sprint(in),
		}}
		arr := &[]*CommentResolver{comment, comment}
		// append the Comments to the Result
		results = append(results, &dataloader.Result{Data: arr})
	}
	return results
}
