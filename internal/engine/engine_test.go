package engine

import (
	"testing"

	"github.com/yyysay/registry-sync/internal/destination"
	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/mapper"
	"github.com/yyysay/registry-sync/internal/rule"
	"github.com/yyysay/registry-sync/internal/source"
)

func TestGenerate(t *testing.T) {
	r := rule.New(
		"ghcr-to-aliyun",
		source.New(
			"ghcr.io",
			source.Default,
		),
		destination.New(
			"registry.example.com/images",
			mapper.New(mapper.Basename),
		),
	)

	engine := New([]*rule.Rule{r})

	img, err := image.Parse(
		"ghcr.io/sagernet/sing-box:latest",
	)
	if err != nil {
		t.Fatal(err)
	}

	tasks, err := engine.Generate(
		[]*image.Image{img},
	)

	if err != nil {
		t.Fatalf(
			"Generate() error = %v",
			err,
		)
	}

	if len(tasks) != 1 {
		t.Fatalf(
			"tasks = %d",
			len(tasks),
		)
	}

	want := "registry.example.com/images/sing-box:latest"

	if tasks[0].Target.Reference != want {
		t.Fatalf(
			"Target = %q, want %q",
			tasks[0].Target.Reference,
			want,
		)
	}
}

func TestGenerateNoMatch(t *testing.T) {
	r := rule.New(
		"docker-only",
		source.New(
			"docker.io",
			source.Default,
		),
		destination.New(
			"registry.example.com/images",
			mapper.New(mapper.Basename),
		),
	)

	engine := New([]*rule.Rule{r})

	img, err := image.Parse(
		"ghcr.io/sagernet/sing-box:latest",
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = engine.Generate(
		[]*image.Image{img},
	)

	if err == nil {
		t.Fatal("expected error")
	}
}
