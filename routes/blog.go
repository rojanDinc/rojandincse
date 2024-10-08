package routes

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type BlogPage struct {
	Posts []string
}

func BlogHandler() http.Handler {
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
			Posts: make([]string, 0, len(posts)),
		}

		for _, post := range posts {
			s := strings.Split(post, ".")
			blogPage.Posts = append(blogPage.Posts, s[0])
		}

		temp, err := template.ParseFiles(filepath.Join("templates", "blog.html"))
		if err != nil {
			log.Println("failed to parse template: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := temp.Execute(w, blogPage); err != nil {
			log.Println("failed to execute template: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
