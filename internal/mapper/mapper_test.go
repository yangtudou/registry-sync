package mapper

import (
	"testing"

	"github.com/yyysay/registry-sync/internal/image"
)

func TestMapPreserve(t *testing.T) {
	img, err := image.Parse(
		"ghcr.io/cloudflare/cloudflared:latest",
	)
	if err != nil {
		t.Fatal(err)
	}

	dst := New(Preserve).Map(img)

	if dst.Name != "cloudflare/cloudflared" {
		t.Fatalf(
			"Name = %q",
			dst.Name,
		)
	}

	if dst.Registry != "ghcr.io" {
		t.Fatalf(
			"Registry = %q",
			dst.Registry,
		)
	}
}

func TestMapBasename(t *testing.T) {
	img, err := image.Parse(
		"cloudflare/cloudflared:latest",
	)
	if err != nil {
		t.Fatal(err)
	}

	dst := New(Basename).Map(img)

	if dst.Name != "cloudflared" {
		t.Fatalf(
			"Name = %q",
			dst.Name,
		)
	}
}

func TestMapGHCRBasename(t *testing.T) {
	img, err := image.Parse(
		"ghcr.io/sagernet/sing-box:latest",
	)
	if err != nil {
		t.Fatal(err)
	}

	dst := New(Basename).Map(img)

	if dst.Name != "sing-box" {
		t.Fatalf(
			"Name = %q",
			dst.Name,
		)
	}
}
