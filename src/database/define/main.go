package main

import (
	"gorm.io/gen"
)

func NewGenerator() *gen.Generator {
	g := gen.NewGenerator(gen.Config{
		OutPath: "src/query",
	})

	return g
}

func main() {
	g := NewGenerator()
	generateMoney(g)
	generateUser(g)
	g.Execute()
}
