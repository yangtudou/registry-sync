package mapper

import (
	"github.com/yyysay/registry-sync/internal/image"
)

type Mode string

const (
	Basename Mode = "basename"
	Preserve Mode = "preserve"
)

type Mapper struct {
	mode Mode
}

func New(mode Mode) *Mapper {
	return &Mapper{
		mode: mode,
	}
}

func (m *Mapper) Map(img *image.Image) *image.Image {
	dst := &image.Image{
		Registry: img.Registry,
		Tag:      img.Tag,
	}

	switch m.mode {
	case Preserve:
		dst.Name = img.Name
		dst.Namespace = img.Namespace

	case Basename:
		dst.Name = basename(img.Name)

	default:
		dst.Name = img.Name
		dst.Namespace = img.Namespace
	}

	return dst
}

func basename(name string) string {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' {
			return name[i+1:]
		}
	}

	return name
}
