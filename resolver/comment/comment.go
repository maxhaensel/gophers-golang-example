package commentresolver

import (
	"context"
	"fmt"
	"graphql/model"

	"github.com/graph-gophers/dataloader"
)

// ResponseComment to resolve the Comment-Type
type ResponseComment struct {
	comment *model.Comment
}

// UID ResponseComment
func (r *ResponseComment) UID(ctx context.Context) *string {
	return &r.comment.UID
}

// Name ResponseComment
func (r *ResponseComment) Name(ctx context.Context) *string {
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
		comment := &ResponseComment{comment: &model.Comment{
			UID:  key.String(),
			Name: "index: " + fmt.Sprint(in),
		}}
		arr := &[]*ResponseComment{comment, comment}
		// append the Comments to the Result
		results = append(results, &dataloader.Result{Data: arr})
	}
	return results
}
