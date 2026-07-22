package image

import (
	"bufio"
	"os"
	"strings"
)

func Load(filename string) ([]*Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var images []*Image

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		img, err := Parse(line)
		if err != nil {
			return nil, err
		}

		images = append(images, img)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return images, nil
}
