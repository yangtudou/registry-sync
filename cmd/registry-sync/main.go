package main

import (
	"fmt"
	"log"

	"github.com/yyysay/registry-sync/internal/destination"
	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/mapper"
	"github.com/yyysay/registry-sync/internal/task"
)

func main() {
	images, err := image.Load("images.txt")
	if err != nil {
		log.Fatal(err)
	}

	dst := destination.New(
		"registry.example.com/myspace",
		mapper.New(mapper.Basename),
	)

	tasks := task.Generate(images, dst)

	for _, t := range tasks {
		fmt.Printf(
			"%s -> %s\n",
			t.Source.Reference,
			t.Target.Reference,
		)
	}
}
