package engine

import (
	"context"

	"registry-sync/model"
)

func (e *Engine) Execute(
	ctx context.Context,
	plan model.Plan,
) error {

	sources := ResolveSources(plan)

	for _, target := range plan.Targets {

		targetImage := BuildTargetImage(
			plan.Image,
			target,
		)

		var lastErr error

		for _, source := range sources {

			dumpCopyTask(
				source,
				targetImage,
				plan.Image.Platform,
			)

			err := e.copier.Copy(
				ctx,
				source,
				targetImage,
				plan.Image.Platform,
			)

			if err == nil {
				lastErr = nil
				break
			}

			lastErr = err
		}

		if lastErr != nil {
			return lastErr
		}
	}

	return nil
}
