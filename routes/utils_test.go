package routes

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractFrontmatter(t *testing.T) {
	content := `---
title: test
published_at: 2025-01-01
---
# Hello world!
`
	file, err := os.CreateTemp(os.TempDir(), "*.md")
	require.NoError(t, err)
	filePath := file.Name()

	_, err = io.WriteString(file, content)
	require.NoError(t, err)
	require.NoError(t, file.Close())

	fm, err := extractFrontmatter(filePath)
	require.NoError(t, err)

	assert.Equal(t, "2025-01-01", fm.PublishedAt)
	assert.Equal(t, "test", fm.Title)
}
