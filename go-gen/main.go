package main

import (
	"fmt"
)

//go:generate go run github.com/blaqkube/blaqkode/go-gen/gen

type Demo struct {
	Field1 string
}

var demo = Demo{
	Field1: "Hello, World",
}

func main() {
	fmt.Println(demo.Field1)
}
