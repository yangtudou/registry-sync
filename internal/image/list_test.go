package image

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()

	file := filepath.Join(dir, "images.txt")

	content := `# docker
nginx

cloudflare/cloudflared

# github
ghcr.io/sagernet/sing-box
`

	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	images, err := Load(file)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if len(images) != 3 {
		t.Fatalf("len(images) = %d, want 3", len(images))
	}

	if images[0].Reference != "nginx:latest" {
		t.Fatalf("images[0].Reference = %q", images[0].Reference)
	}

	if images[1].Reference != "cloudflare/cloudflared:latest" {
		t.Fatalf("images[1].Reference = %q", images[1].Reference)
	}

	if images[2].Reference != "ghcr.io/sagernet/sing-box:latest" {
		t.Fatalf("images[2].Reference = %q", images[2].Reference)
	}
}
