package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mutahirlatif/go-clean-architecture/task"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc task.UseCase) {
	h := NewHandler(uc)

	bookmarks := router.Group("/tasks")
	{
		bookmarks.POST("", h.Create)
		bookmarks.GET("", h.Get)
		bookmarks.DELETE("", h.Delete)
	}
}
