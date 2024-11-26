package api

import "net/http"

type api struct {
}

func (a *api) getEvents(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func (a *api) postGitHubWebhook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func New() http.Handler {
	api := api{}
	rootMux := http.NewServeMux()

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("GET /events", api.getEvents)
	rootMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	githubMux := http.NewServeMux()
	githubMux.HandleFunc("POST /webhook", api.postGitHubWebhook)
	rootMux.Handle("/github/", http.StripPrefix("/github", githubMux))

	return rootMux
}
