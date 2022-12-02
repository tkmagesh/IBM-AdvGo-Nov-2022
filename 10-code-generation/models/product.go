package models

/*
go:generate go run ../col-gen.go -N Product -P models
go:generate go fmt Products.go
*/

//go:generate col-gen-util -N Product -P models
//go:generate go fmt Products.go
type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}
