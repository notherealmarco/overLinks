package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/api", rt.getHelloWorld)
	rt.router.GET("/api/links", rt.wrap(rt.getLinks))
	rt.router.POST("/api/links", rt.wrap(rt.postLink))
	rt.router.DELETE("/api/links/:id", rt.wrap(rt.deleteLink))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
