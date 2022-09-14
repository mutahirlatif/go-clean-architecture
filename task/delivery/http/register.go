package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mutahirlatif/go-clean-architecture/task"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc task.UseCase) {
	h := NewHandler(uc)

	tasks := router.Group("/tasks")
	{
		tasks.POST("", h.Create)
		tasks.GET("", h.Get)
		tasks.DELETE("", h.Delete)
		tasks.PUT("", h.Put)
	}
}
