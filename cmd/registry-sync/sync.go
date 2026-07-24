package main

import (
	"context"
	"fmt"

	"registry-sync/config"
	"registry-sync/copier"
	"registry-sync/engine"
	"registry-sync/planner"
)

func SyncCommand(
	args []string,
) error {

	if len(args) < 1 {

		return fmt.Errorf(
			"usage: registry-sync sync <config.yaml>",
		)
	}

	cfg, err := config.LoadConfig(
		args[0],
	)

	if err != nil {
		return err
	}

	plans := planner.Build(cfg)

	fmt.Println(
		planner.Dump(plans),
	)

	e := engine.New(
		copier.New(),
	)

	ctx := context.Background()

	for _, plan := range plans {

		err := e.Execute(
			ctx,
			plan,
		)

		if err != nil {

			return fmt.Errorf(
				"sync failed: %w",
				err,
			)
		}
	}

	return nil
}
