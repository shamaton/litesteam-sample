//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate("../../ent/schema", &gen.Config{
		Target: "../",
	}); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
