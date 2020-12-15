package main

import (
	"graphql/resolver"
	"graphql/schema"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	// Fetch root-schema
	s := schema.GetRootSchema()
	// parse schema and validate
	schema := graphql.MustParseSchema(*s, &resolver.Resolver{})
	// create basic http-handler
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
