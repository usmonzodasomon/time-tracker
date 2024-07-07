package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/usmonzodasomon/time-tracker/internal/model"
	"github.com/usmonzodasomon/time-tracker/internal/repository"
	"github.com/usmonzodasomon/time-tracker/internal/service"
	"github.com/usmonzodasomon/time-tracker/pkg/logger"
	"log/slog"
	"net/http"
	"strconv"
)

type taskHandler struct {
	service     service.TaskServiceI
	userService service.UserServiceI
}

func newTaskHandler(handler *gin.RouterGroup, db *sqlx.DB) {
	taskRepo := repository.NewTaskRepo(db)
	userRepo := repository.NewUserRepo(db)

	taskService := service.NewTaskService(taskRepo)
	userService := service.NewUserService(userRepo)

	r := &taskHandler{
		service:     taskService,
		userService: userService,
	}

	h := handler.Group("/task")
	{
		h.POST("/", r.CreateTask)
		h.POST("/:task_id/start", r.StartTask)
		h.POST("/:task_id/stop", r.StopTask)
	}
}

// CreateTask add a new task.
//
// @Summary Create a new task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param request body model.TaskRequestBody true "Task details"
// @Success 201 {object} SuccessResponse  "Task ID"
// @Failure 400 {object} ErrorResponse "Error message"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /task [post]
func (h *taskHandler) CreateTask(c *gin.Context) {
	logger.Logger.Info("start create task")
	var input model.TaskRequestBody

	if err := c.BindJSON(&input); err != nil {
		logger.Logger.Error("error binding json", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("error binding json"))
		return
	}

	logger.Logger.Debug("parsed input", slog.Any("input", input))
	_, err := h.userService.GetUser(input.UserID)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			logger.Logger.Warn("user not found", slog.Int("user_id", input.UserID))
			c.JSON(http.StatusBadRequest, newErrorResponse("user not found"))
			return
		}
		logger.Logger.Error("error getting user", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, newErrorResponse("error getting user"))
		return
	}

	taskID, err := h.service.CreateTask(model.Task{
		UserID:      input.UserID,
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		logger.Logger.Error("error creating task", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, newErrorResponse("error creating task"))
		return
	}

	logger.Logger.Info("task created")
	logger.Logger.Debug(fmt.Sprintf("task with id %d created", taskID))
	c.JSON(http.StatusCreated, newSuccessResponse(strconv.Itoa(taskID)))
}

// StartTask starts a task.
// @Summary Start a task
// @Tags Tasks
// @Produce json
// @Param task_id path int true "Task ID"
// @Success 200 {object} SuccessResponse "Status of the task"
// @Failure 400 {object} ErrorResponse "Error message"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /task/{task_id}/start [post]
func (h *taskHandler) StartTask(c *gin.Context) {
	logger.Logger.Info("start task")
	taskID, err := h.getTaskID(c)
	if err != nil {
		logger.Logger.Error("error getting task id from param", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("error getting task id from param"))
		return
	}

	_, err = h.service.GetTask(taskID)
	if err != nil {
		if errors.Is(err, model.ErrTaskNotFound) {
			logger.Logger.Warn("task not found", slog.Int("task_id", taskID))
			c.JSON(http.StatusBadRequest, newErrorResponse("task not found"))
			return
		}
		logger.Logger.Error("error getting task", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, newErrorResponse("error getting task"))
		return
	}

	err = h.service.StartTask(taskID)
	if err != nil {
		if errors.Is(err, model.ErrTaskAlreadyStarted) {
			logger.Logger.Warn("task already started", slog.Int("task_id", taskID))
			c.JSON(http.StatusBadRequest, newErrorResponse("task already started"))
			return
		}
		logger.Logger.Error("error starting task", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, newErrorResponse("error starting task"))
		return
	}

	logger.Logger.Info("task started")
	logger.Logger.Debug(fmt.Sprintf("task with id %d started", taskID))
	c.JSON(http.StatusOK, newSuccessResponse("task started"))
}

// StopTask stops a task.
// @Summary Stop a task
// @Tags Tasks
// @Produce json
// @Param task_id path int true "Task ID"
// @Success 200 {object} SuccessResponse "Status of the task"
// @Failure 400 {object} ErrorResponse "Error message"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /task/{task_id}/stop [post]
func (h *taskHandler) StopTask(c *gin.Context) {
	logger.Logger.Info("stop task")
	taskID, err := h.getTaskID(c)
	if err != nil {
		logger.Logger.Error("error getting task id from param", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("error getting task id from param"))
		return
	}

	_, err = h.service.GetTask(taskID)
	if err != nil {
		if errors.Is(err, model.ErrTaskNotFound) {
			logger.Logger.Warn("task not found", slog.Int("task_id", taskID))
			c.JSON(http.StatusBadRequest, newErrorResponse("task not found"))
			return
		}
		logger.Logger.Error("error getting task", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, newErrorResponse("error getting task"))
		return
	}

	err = h.service.StopTask(taskID)
	if err != nil {
		if errors.Is(err, model.ErrTaskAlreadyStopped) {
			logger.Logger.Warn("task already stopped", slog.Int("task_id", taskID))
			c.JSON(http.StatusBadRequest, newErrorResponse("task already stopped"))
			return
		}
		logger.Logger.Error("error stopping task", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Logger.Info("task stopped")
	logger.Logger.Debug(fmt.Sprintf("task with id %d stopped", taskID))
	c.JSON(http.StatusOK, gin.H{"status": "task stopped"})
}

func (h *taskHandler) getTaskID(c *gin.Context) (int, error) {
	id := c.Param("task_id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return taskID, nil
}
