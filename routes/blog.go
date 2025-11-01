package routes

import (
	"html/template"
	"log/slog"
	"net/http"
	"path/filepath"
	"strings"
)

type FrontMatter struct {
	Title       string `json:"title"`
	PublishedAt string `json:"published_at"`
}

type Post struct {
	FrontMatter
	Link string
}

type BlogPage struct {
	PageMeta
	Posts []Post
}

func BlogHandler(template *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		postPaths, err := filepath.Glob(filepath.Join("posts", "*.md"))
		if err != nil {
			slog.Error("failed to read posts", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		blogPage := BlogPage{
			PageMeta: PageMeta{
				Title:       "Blog",
				Description: "Blog posts",
				Keywords:    "blog, posts",
			},
			Posts: make([]Post, 0, len(postPaths)),
		}

		for _, postPath := range postPaths {
			link := strings.TrimSuffix(filepath.Base(postPath), ".md")
			fm, err := extractFrontmatter(postPath)
			if err != nil {
				slog.Error("failed to extract frontmatter", "err", err.Error(), "file", postPath)
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			blogPage.Posts = append(blogPage.Posts, Post{
				FrontMatter: *fm,
				Link:        link,
			})
		}

		if err := template.ExecuteTemplate(w, "blog.html", blogPage); err != nil {
			slog.Error("failed to execute template", "error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	})
}
