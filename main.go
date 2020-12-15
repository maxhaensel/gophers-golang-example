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
	s := schema.GetRootSchema()
	schema := graphql.MustParseSchema(*s, &resolver.Resolver{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
