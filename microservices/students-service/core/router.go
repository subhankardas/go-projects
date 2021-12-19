package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	mRouter *mux.Router
}

type RouterConfig struct {
	StrictSlash bool
}

func (router *Router) Init(config RouterConfig) {
	router.mRouter = mux.NewRouter().StrictSlash(config.StrictSlash)
}

func (router *Router) HandleFunc(path string, handler func(http.ResponseWriter,
	*http.Request)) *mux.Route {
	return router.mRouter.HandleFunc(path, handler)
}

func (router *Router) HttpListenAndServe() {
	http.ListenAndServe(":8080", router.mRouter)
}
