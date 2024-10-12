package routes

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type BlogPage struct {
	PageMeta
	Posts []string
}

func BlogHandler(template *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		postPaths, err := filepath.Glob(filepath.Join("posts", "*.md"))
		if err != nil {
			log.Println("failed to read posts: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		posts := make([]string, 0, len(postPaths))
		for _, post := range postPaths {
			posts = append(posts, filepath.Base(post))
		}

		blogPage := BlogPage{
			PageMeta: PageMeta{
				Title:       "Blog",
				Description: "Blog posts",
				Keywords:    "blog, posts",
			},
			Posts: make([]string, 0, len(posts)),
		}

		for _, post := range posts {
			s := strings.Split(post, ".")
			blogPage.Posts = append(blogPage.Posts, s[0])
		}

		if err := template.ExecuteTemplate(w, "blog.html", blogPage); err != nil {
			log.Println("failed to execute template: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
