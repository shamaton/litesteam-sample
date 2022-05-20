//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate("../schema", &gen.Config{
		Target:  "../../db/",
		Package: "github.com/shamaton/litestream-sample/db",
	}); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
