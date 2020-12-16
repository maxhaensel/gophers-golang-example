package main

import (
	rootresolver "graphql/resolver"
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
	schema := graphql.MustParseSchema(*s, &rootresolver.Resolver{})
	// create basic http-handler

	// schema.Exec()

	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
