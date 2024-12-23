package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

type PostPage struct {
	PageMeta
	Content template.HTML
}

func PostHandler(t *template.Template) http.Handler {
	gm := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.DefinitionList,
			extension.Strikethrough,
			extension.Table,
			extension.Linkify,
			extension.TaskList,
		),
		goldmark.WithRendererOptions(
			html.WithWriter(html.NewWriter()),
		),
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		postName := r.PathValue("post")
		if postName == "" {
			http.NotFound(w, r)
			return
		}

		postPath := filepath.Join("posts", postName+".md")

		md, err := markdownToHTML(gm, postPath)
		if err != nil {
			slog.Error("failed to convert markdown to HTML", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		postPage := PostPage{
			PageMeta: PageMeta{
				Title:       fmt.Sprintf("Post - %s", postName),
				Description: fmt.Sprintf("A post where Rojan writes about %s", postName),
				Keywords:    postName,
			},
			Content: template.HTML(md),
		}

		if err := t.ExecuteTemplate(w, "post.html", postPage); err != nil {
			slog.Error("failed to execute template", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func markdownToHTML(gm goldmark.Markdown, path string) (string, error) {
	var buf bytes.Buffer
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	if err := gm.Convert(b, &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}
