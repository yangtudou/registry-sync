package copier

import (
	"context"

	"registry-sync/engine"

	"github.com/google/go-containerregistry/pkg/crane"
)

type CraneCopier struct {
	copyFunc func(
		string,
		string,
		...crane.Option,
	) error
}

func New() *CraneCopier {

	return &CraneCopier{
		copyFunc: crane.Copy,
	}
}

func (c *CraneCopier) Copy(
	ctx context.Context,
	source string,
	target string,
	platform []string,
) error {

	opts := []crane.Option{
		crane.WithContext(ctx),
	}

	opts = append(
		opts,
		buildPlatformOption(platform)...,
	)

	err := c.copyFunc(
		source,
		target,
		opts...,
	)

	if err != nil {

		dumpFailed(err)

		return err
	}

	dumpSuccess()

	return nil
}

var _ engine.Copier = (*CraneCopier)(nil)
