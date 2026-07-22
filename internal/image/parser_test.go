package image

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		registry  string
		namespace string
		imageName string
		tag       string
		reference string
	}{
		{
			name:      "docker official",
			input:     "nginx",
			registry:  "docker.io",
			namespace: "library",
			imageName: "library/nginx",
			tag:       "latest",
			reference: "nginx:latest",
		},
		{
			name:      "docker namespace",
			input:     "cloudflare/cloudflared",
			registry:  "docker.io",
			namespace: "cloudflare",
			imageName: "cloudflare/cloudflared",
			tag:       "latest",
			reference: "cloudflare/cloudflared:latest",
		},
		{
			name:      "ghcr",
			input:     "ghcr.io/sagernet/sing-box",
			registry:  "ghcr.io",
			namespace: "sagernet",
			imageName: "sagernet/sing-box",
			tag:       "latest",
			reference: "ghcr.io/sagernet/sing-box:latest",
		},
		{
			name:      "with tag",
			input:     "registry.k8s.io/pause:3.10",
			registry:  "registry.k8s.io",
			namespace: "",
			imageName: "pause",
			tag:       "3.10",
			reference: "registry.k8s.io/pause:3.10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := Parse(tt.input)
			if err != nil {
				t.Fatal(err)
			}

			if img.Registry != tt.registry {
				t.Fatalf(
					"Registry = %q, want %q",
					img.Registry,
					tt.registry,
				)
			}

			if img.Namespace != tt.namespace {
				t.Fatalf(
					"Namespace = %q, want %q",
					img.Namespace,
					tt.namespace,
				)
			}

			if img.Name != tt.imageName {
				t.Fatalf(
					"Name = %q, want %q",
					img.Name,
					tt.imageName,
				)
			}

			if img.Tag != tt.tag {
				t.Fatalf(
					"Tag = %q, want %q",
					img.Tag,
					tt.tag,
				)
			}

			if img.Reference != tt.reference {
				t.Fatalf(
					"Reference = %q, want %q",
					img.Reference,
					tt.reference,
				)
			}
		})
	}
}
