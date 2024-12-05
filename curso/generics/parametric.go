package main

// PARAMETRIC CONSTRAINTS
// Parametric constraints allow us to restrict the types that can be used with a generic function. We can use constraints to ensure that a generic
// function only works with types that have certain methods.

// EXAMPLE

// The store interface represents a store that sells products.
// It takes a type parameter P that must implement the product interface.
type store[P product] interface {
	Sell(P) // A store is anything that can sell a product
}

type product interface {
	Price() float64
	Name() string
}
