package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/frontmatter"
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
			&frontmatter.Extender{},
			highlighting.NewHighlighting(
				highlighting.WithStyle("github-dark"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
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

		fm, md, err := markdownToHTML(gm, postPath)
		if err != nil {
			slog.Error("failed to convert markdown to HTML", "error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		postPage := PostPage{
			PageMeta: PageMeta{
				Title:       fmt.Sprintf("Post - %s", postName),
				Description: fmt.Sprintf("A post where Rojan writes about %s", postName),
				Keywords:    postName,
				FrontMatter: fm,
			},
			Content: template.HTML(md),
		}

		if err := t.ExecuteTemplate(w, "post.html", postPage); err != nil {
			slog.Error("failed to execute template", "error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	})
}

func markdownToHTML(gm goldmark.Markdown, path string) (FrontMatter, string, error) {
	var buf bytes.Buffer
	var fm FrontMatter
	b, err := os.ReadFile(path)
	if err != nil {
		return fm, "", err
	}

	ctx := parser.NewContext()
	if err := gm.Convert(b, &buf, parser.WithContext(ctx)); err != nil {
		return fm, "", err
	}

	d := frontmatter.Get(ctx)
	if err := d.Decode(&fm); err != nil {
		return fm, "", err
	}

	return fm, buf.String(), nil
}
