package routes

import "net/http"

type (
	githubHandler struct{}
)

// ConfigureGithubRoutes for the server
func ConfigureGithubRoutes(mux *http.ServeMux) {
	handler := githubHandler{}
	mux.Handle("/github", &handler)
}

func (h *githubHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
}
