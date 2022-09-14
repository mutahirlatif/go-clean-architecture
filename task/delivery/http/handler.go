package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mutahirlatif/go-clean-architecture/auth"
	"github.com/mutahirlatif/go-clean-architecture/models"
	"github.com/mutahirlatif/go-clean-architecture/task"
)

type Task struct {
	ID         string    `json:"id"`
	TaskDetail string    `json:"taskDetail"`
	DueDate    time.Time `json:"dueDate"`
}

type Handler struct {
	useCase task.UseCase
}

func NewHandler(useCase task.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInput struct {
	TaskDetail string    `json:"taskDetail"`
	DueDate    time.Time `json:"dueDate"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreateTask(c.Request.Context(), user, inp.TaskDetail, inp.DueDate); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type getResponse struct {
	Tasks []*Task `json:"tasks"`
}

func (h *Handler) Get(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	tms, err := h.useCase.GetTasks(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getResponse{
		Tasks: toTasks(tms),
	})
}

type updateInput struct {
	ID         string    `json:"id"`
	TaskDetail string    `json:"taskDetail"`
	DueDate    time.Time `json:"dueDate"`
}

func (h *Handler) Put(c *gin.Context) {
	inp := new(updateInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)
	if _, err := h.useCase.GetTaskByID(c.Request.Context(), user, inp.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := h.useCase.UpdateTask(c.Request.Context(), user, inp.TaskDetail, inp.DueDate, inp.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type deleteInput struct {
	ID string `json:"id"`
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(deleteInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.DeleteTask(c.Request.Context(), user, inp.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func toTasks(ts []*models.Task) []*Task {
	out := make([]*Task, len(ts))

	for i, b := range ts {
		out[i] = toTask(b)
	}

	return out
}

func toTask(t *models.Task) *Task {
	return &Task{
		ID:         t.ID,
		TaskDetail: t.TaskDetail,
		DueDate:    t.DueDate,
	}
}
