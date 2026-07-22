package engine

import (
	"fmt"

	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/rule"
	"github.com/yyysay/registry-sync/internal/task"
)

type Engine struct {
	Rules []*rule.Rule
}

func New(rules []*rule.Rule) *Engine {
	return &Engine{
		Rules: rules,
	}
}

func (e *Engine) Generate(images []*image.Image) ([]*task.Task, error) {
	var tasks []*task.Task

	for _, img := range images {
		r, err := e.match(img)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task.Task{
			Source: img,
			Target: r.Destination.Map(img),
		})
	}

	return tasks, nil
}

func (e *Engine) match(img *image.Image) (*rule.Rule, error) {
	for _, r := range e.Rules {
		if r.Source.Registry == img.Registry {
			return r, nil
		}
	}

	return nil, fmt.Errorf(
		"no rule matched image: %s",
		img.Reference,
	)
}
