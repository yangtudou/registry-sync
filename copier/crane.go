package copier

import (
	"context"
	"fmt"

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

	fmt.Println("CRANE COPY")
	fmt.Println("==========")

	fmt.Println("SOURCE:")
	fmt.Println(" ", source)

	fmt.Println()

	fmt.Println("TARGET:")
	fmt.Println(" ", target)

	fmt.Println()

	opts := []crane.Option{
		crane.WithContext(ctx),
	}

	if len(platform) > 0 {

		fmt.Println("PLATFORM:")

		for _, p := range platform {
			fmt.Println(" ", p)
		}

		fmt.Println()

		opts = append(
			opts,
			buildPlatformOption(platform)...,
		)
	}

	fmt.Println("ENGINE:")
	fmt.Println(" crane.Copy")

	fmt.Println()

	err := c.copyFunc(
		source,
		target,
		opts...,
	)

	if err != nil {

		fmt.Println("STATUS:")
		fmt.Println(" FAILED")

		fmt.Println("ERROR:")
		fmt.Println(" ", err)

		fmt.Println()

		return err
	}

	fmt.Println("STATUS:")
	fmt.Println(" SUCCESS")
	fmt.Println()

	return nil
}

var _ engine.Copier = (*CraneCopier)(nil)
