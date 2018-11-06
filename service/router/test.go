package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//RegisterTestRouter register test Apis
func (r *Router) RegisterTestRouter() {
	test := r.E.Group("/test")
	{
		test.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "What happend?")
		})
	}
}
