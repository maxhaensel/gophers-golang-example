package userresolver

import (
	"context"
	"graphql/model"

	"github.com/graph-gophers/graphql-go"
)

// ResolverUser export
type ResolverUser struct{}

// ResponseUser ...
type ResponseUser struct {
	User *model.User
}

// User Resolver
func (r *ResolverUser) User(ctx context.Context, args struct {
	UID string
}) (*ResponseUser, error) {
	user := &model.User{UID: "id1", Name: "my Name"}
	return &ResponseUser{User: user}, nil
}

// UID resolver
func (r *ResponseUser) UID() *graphql.ID {
	uid := graphql.ID(r.User.UID)
	return &uid
}

// Name resolver
func (r *ResponseUser) Name() *string {
	return &r.User.Name
}
