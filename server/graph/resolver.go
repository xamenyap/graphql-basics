package graph

import "github.com/xamenyap/graphql-basics/server/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	products []model.Product
}

func NewResolver(products []model.Product) *Resolver {
	return &Resolver{products: products}
}
