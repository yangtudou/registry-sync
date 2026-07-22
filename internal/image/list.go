package image

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Load 从文本文件加载镜像列表。
//
// 规则：
//   - 空行忽略
//   - 以 # 开头的行为注释
//   - 每行一个镜像
func Load(filename string) ([]*Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open %q: %w", filename, err)
	}
	defer file.Close()

	var images []*Image

	scanner := bufio.NewScanner(file)
	line := 0

	for scanner.Scan() {
		line++

		text := strings.TrimSpace(scanner.Text())

		if text == "" {
			continue
		}

		if strings.HasPrefix(text, "#") {
			continue
		}

		img, err := Parse(text)
		if err != nil {
			return nil, fmt.Errorf("%s:%d: %w", filename, line, err)
		}

		images = append(images, img)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read %q: %w", filename, err)
	}

	return images, nil
}
