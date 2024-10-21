package main

import (
	"log/slog"
	"net/http"
	"os"
	"rojandincse/routes"
)

var (
	port = env("PORT", "8080")
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	routes := routes.NewRoutes()
	http.ListenAndServe(":"+port, routes)
}

func env(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
