package router

import (
	"api-gateway/handlers"
	"api-gateway/middlewares"
	"api-gateway/utils"
	"fmt"

	// "fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() http.Handler {
	r := httprouter.New()

	// Service prefix to env key mapping adding all the microservices ports present in .env file
	services := map[string]string{
		"/api/v1/auth/":   utils.Env("AUTH_SERVICE_URL", "http://localhost:4000"),
		"/api/v1/posts/":  utils.Env("POST_SERVICE_URL", "http://localhost:4001"),
		"/api/v1/stories/": utils.Env("STORY_SERVICE_URL", "http://localhost:4002"),
	}

	//all the rest api methods
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
	}

	for prefix, target := range services {
		proxy := handlers.NewReverseProxy(target)
		fmt.Println("proxy", proxy)
		for _, method := range methods {
			//r.handler(method, path, proxy)
			r.Handler(method, prefix+"*path", proxy)
		}
	}
	return middleware.CORSMiddleware(r)
}
