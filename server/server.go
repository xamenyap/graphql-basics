package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/xamenyap/graphql-basics/server/graph"
	"github.com/xamenyap/graphql-basics/server/graph/model"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: graph.NewResolver([]model.Product{
				{
					ID:   "1",
					Name: "macbook pro 2023",
					Reviews: []*model.Review{
						{
							ID:        "1",
							Content:   "very good laptop",
							Rating:    10,
							ProductID: "1",
						},
						{
							ID:        "2",
							Content:   "nice to have",
							Rating:    9,
							ProductID: "1",
						},
					},
				},
				{
					ID:   "2",
					Name: "macbook air 2023",
					Reviews: []*model.Review{
						{
							ID:        "1",
							Content:   "battery could be better",
							Rating:    7,
							ProductID: "2",
						},
						{
							ID:        "2",
							Content:   "should have bought the pro version",
							Rating:    5,
							ProductID: "2",
						},
					},
				},
			})},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
