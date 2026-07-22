package image

// Image 表示一个规范化后的镜像引用。
type Image struct {
	// 用户原始输入
	Raw string

	// 唯一规范引用，例如：
	// docker.io/library/nginx:latest
	Reference string

	// 仓库地址，例如：
	// docker.io
	// ghcr.io
	Registry string

	// namespace/repository，例如：
	// library/nginx
	// cloudflare/cloudflared
	Name string

	// latest
	Tag string

	// sha256:...
	Digest string
}
