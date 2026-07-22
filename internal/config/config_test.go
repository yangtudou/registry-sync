package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()

	file := filepath.Join(dir, "config.yaml")

	content := `
destination:
  registry: registry.cn-hangzhou.aliyuncs.com/myspace
  mode: basename
`

	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(file)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Destination.Registry != "registry.cn-hangzhou.aliyuncs.com/myspace" {
		t.Fatalf(
			"Registry = %q",
			cfg.Destination.Registry,
		)
	}

	if cfg.Destination.Mode != "basename" {
		t.Fatalf(
			"Mode = %q",
			cfg.Destination.Mode,
		)
	}
}
