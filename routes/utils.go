package routes

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strings"

	"sigs.k8s.io/yaml"
)

func extractFrontmatter(filePath string) (*FrontMatter, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var frontmatter strings.Builder
	var inFrontmatter bool
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if lineNum == 1 {
			if line == "---" {
				inFrontmatter = true
				continue
			} else {
				return nil, errors.New("no frontmatter found")
			}
		}

		if inFrontmatter && line == "---" {
			break
		}

		if inFrontmatter {
			if frontmatter.Len() > 0 {
				frontmatter.WriteString("\n")
			}
			frontmatter.WriteString(line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if frontmatter.Len() == 0 {
		return nil, errors.New("no frontmatter")
	}

	var fm FrontMatter
	var buf bytes.Buffer
	buf.WriteString(frontmatter.String())

	if err := yaml.Unmarshal(buf.Bytes(), &fm); err != nil {
		return nil, err
	}

	return &fm, nil
}
