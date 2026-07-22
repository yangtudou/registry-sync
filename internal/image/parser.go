package image

import (
	"fmt"

	"github.com/distribution/reference"
)

// Parse 解析一个 Docker Image Reference，并返回规范化后的 Image。
func Parse(raw string) (*Image, error) {
	named, err := reference.ParseNormalizedNamed(raw)
	if err != nil {
		return nil, fmt.Errorf("parse image %q: %w", raw, err)
	}

	// 自动补充 latest
	named = reference.TagNameOnly(named)

	img := &Image{
		Raw:       raw,
		Reference: reference.FamiliarString(named),
		Registry:  reference.Domain(named),
		Name:      reference.Path(named),
	}

	if tagged, ok := named.(reference.Tagged); ok {
		img.Tag = tagged.Tag()
	}

	if digested, ok := named.(reference.Digested); ok {
		img.Digest = digested.Digest().String()
	}

	return img, nil
}
