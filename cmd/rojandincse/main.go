package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"rojandincse/routes"

	"github.com/Masterminds/sprig/v3"
)

var (
	port = env("PORT", "8080")
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	temp, err := template.New("base").Funcs(sprig.FuncMap()).ParseGlob("templates/*.html")
	if err != nil {
		slog.Error("something went wrong", "err", err.Error())
	}

	routes := routes.NewRoutes(temp)
	slog.Info(fmt.Sprintf("started server on port: %s", port))
	http.ListenAndServe(":"+port, routes)
}

func env(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
