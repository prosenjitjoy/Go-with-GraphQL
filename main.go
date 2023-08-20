package main

import (
	"log"
	"main/schema"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	fileContent, err := os.ReadFile("./schema/schema.gql")
	if err != nil {
		log.Fatal(err)
	}

	query := string(fileContent)
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema, err := graphql.ParseSchema(query, &schema.Resolver{}, opts...)
	if err != nil {
		log.Fatal(err)
	}

	graphql := &relay.Handler{Schema: schema}
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Post("/query", graphql.ServeHTTP)

	err = http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatal(err)
	}
}
