# gophers-golang-example

## Run it Local

```bash
go get -v all
go mod vendor
go generate
go run main.go
```

## How to create new Querys, Mutations or Resolvers

1. Extend the .graphql-schema-files

   ```graphql
   type Todo {
     uid: ID
     name: String
     comment(id: String): [Comment]
   }
   ```

   create some querys and mutations

   ```graphql
   type Query {
     todo(uid: ID!): Todo
     todos(uid: ID!): [Todo]
   }
   type Mutation {
     createTodo(todo: String!): Todo
   }
   ```

2. Create a new package e.g. `todoresolver` in the resolver-folder

   ```golang
   package todoresolver
   ```

3. Create a struct and name it like `ResolverTodo`

   ```golang
   type ResolverTodo struct{}
   ```

4. Create a model that represent the results

   ```golang
   package model
    // Todo Model
    type Todo struct {
    UID string `json:"uid"`
    Name string `json:"name"`
    }
   ```

5. Bind model to Response

   ```golang
   type ResponseTodo struct {
   todo *model.Todo
   }
   ```

6. Create the Resolver-Functions that match the model-properties

   each propertie need his own resolver-function

   ```golang
   func (r *ResponseTodo) Name() *string {
   return &r.todo.Name
   }
   ```

7. Creating the resolver for the querys or mutations and bind them

   ```golang
   // Todo Resolver
   func (r *ResolverTodo) Todo(ctx context.Context, args *struct {
      UID string
   }) (*ResponseTodo, error) {
      todo := &ResponseTodo{todo: &model.Todo{Name: "name", UID: "id1"}}
      return todo, nil
   }
   ```

8. Bind resolver to `rootresolver`

   ```golang
   package rootresolver

   import (
       todoresolver "graphql/resolver/todo"
       ...
   )
   // Resolver export
   type Resolver struct {
       todoresolver.ResolverTodo
       ...
   }
   ```

## Build Binary

t.b.a
