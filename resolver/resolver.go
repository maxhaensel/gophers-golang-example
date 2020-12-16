package rootresolver

import (
	todoresolver "graphql/resolver/todo"
	userresolver "graphql/resolver/user"
)

// Resolver export
type Resolver struct {
	todoresolver.ResolverTodo
	userresolver.ResolverUser
}
