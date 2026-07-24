package engine

import (
	"registry-sync/model"
)

func ResolveSources(
	plan model.Plan,
) []string {

	var sources []string

	// mirror 优先
	sources = append(
		sources,
		BuildMirrorImages(
			plan.Image,
			plan.Mirrors,
		)...,
	)

	// 原始源最后
	sources = append(
		sources,
		BuildImageName(
			plan.Image.Registry,
			plan.Image.Repository,
			plan.Image.Tag,
		),
	)

	return sources
}
