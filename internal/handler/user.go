package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/usmonzodasomon/time-tracker/internal/external_api"
	"github.com/usmonzodasomon/time-tracker/internal/external_api/mocks"
	"github.com/usmonzodasomon/time-tracker/internal/model"
	"github.com/usmonzodasomon/time-tracker/internal/repository"
	"github.com/usmonzodasomon/time-tracker/internal/service"
	"github.com/usmonzodasomon/time-tracker/pkg/logger"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type userHandler struct {
	service         service.UserServiceI
	externalApiInfo external_api.UserExternalInfoI
}

func newUserHandler(handler *gin.RouterGroup, db *sqlx.DB) {
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	externalApiInfo := mocks.NewUserExternalInfo()

	r := &userHandler{
		service:         userService,
		externalApiInfo: externalApiInfo,
	}

	h := handler.Group("/user")
	{
		h.GET("/", r.GetAllUsers)
		h.GET("/:user_id/time-spent", r.GetUserTimeSpent)
		h.POST("/", r.CreateUser)
		h.PATCH("/:user_id", r.UpdateUser)
		h.DELETE("/:user_id", r.DeleteUser)
	}
}

// GetAllUsers retrieves all users based on the provided filter.
// @Summary Get all users
// @Tags Users
// @Produce json
// @Param filters query model.UserFilter true "Filters"
// @Success 200 {array} model.User "List of users"
// @Failure 400 {object} ErrorResponse "Error message"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /user [get]
func (h *userHandler) GetAllUsers(c *gin.Context) {
	logger.Logger.Info("start get all users")
	var filter model.UserFilter

	if err := c.BindQuery(&filter); err != nil {
		logger.Logger.Error("error binding query", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("error binding query"))
		return
	}

	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.PerPage == 0 {
		filter.PerPage = 10
	}
	logger.Logger.Debug("parsed filter", slog.Any("filter", filter))
	users, err := h.service.GetAllUsers(filter)
	if err != nil {
		logger.Logger.Error("error getting users", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, newErrorResponse("error getting users"))
		return
	}

	logger.Logger.Info("got users")
	logger.Logger.Debug("got users", slog.Any("users", users))
	c.JSON(http.StatusOK, users)
}

// GetUserTimeSpent retrieves the time spent by the user based on the provided user ID and period.
// @Summary Get user time spent
// @Tags Users
// @Produce json
// @Param user_id path int true "User ID"
// @Param start_period query string true "Start period" example("2023-30-12 00:00:00")
// @Param end_period query string true "End period" example("2023-30-12 23:59:59")
// @Success 200 {array} model.TaskTimeSpent "List of time spent"
// @Failure 400 {object} ErrorResponse "Error message"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /user/{user_id}/time-spent [get]
func (h *userHandler) GetUserTimeSpent(c *gin.Context) {
	logger.Logger.Info("start get user time spent")
	id := c.Param("user_id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		logger.Logger.Error("error parsing user id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("error parsing user id"))
		return
	}

	startPeriod, err := time.Parse("2006-01-02 15:04:05", c.Query("start_period"))
	if err != nil {
		logger.Logger.Error("error parsing start date", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("Invalid start date format"))
		return
	}

	endPeriod, err := time.Parse("2006-01-02 15:04:05", c.Query("end_period"))
	if err != nil {
		logger.Logger.Error("error parsing end date", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("Invalid end date format"))
		return
	}

	logger.Logger.Debug("parsed ",
		slog.Int("user_id", userID),
		slog.Any("start_period", startPeriod),
		slog.Any("end_period", endPeriod))

	timeSpent, err := h.service.GetUserTimeSpent(userID, startPeriod, endPeriod)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			logger.Logger.Warn("user not found", slog.Int("user_id", userID))
			c.JSON(http.StatusBadRequest, newErrorResponse("user not found"))
			return
		}
		logger.Logger.Error("error getting time spent", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, newErrorResponse("error getting time spent"))
		return
	}
	logger.Logger.Info("got time spent")
	logger.Logger.Debug("got time spent", slog.Any("time_spent", timeSpent))
	c.JSON(http.StatusOK, timeSpent)
}

// CreateUser creates a new user
// @Summary Create a user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body model.UserRequestBody true "User details"
// @Success 201 {object} SuccessResponse "User ID"
// @Failure 400 {object} ErrorResponse "Error message"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /user [post]
func (h *userHandler) CreateUser(c *gin.Context) {
	logger.Logger.Info("start create user")
	var input model.UserRequestBody
	if err := c.BindJSON(&input); err != nil {
		logger.Logger.Error("error binding json", slog.String("error", err.Error()))
		c.JSON(http.StatusOK, newErrorResponse("error binding json"))
		return
	}

	passportSerie, passportNumber, err := parsePassportNumberAndSerie(input.PassportNumber)
	if err != nil {
		logger.Logger.Error("error parsing passport number", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("error parsing passport number"))
		return
	}

	user, err := h.externalApiInfo.GetUser(passportSerie, passportNumber)
	if err != nil {
		logger.Logger.Error("error getting user from external api", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, newErrorResponse("error getting user from external api"))
		return
	}
	logger.Logger.Debug("got user from external api")

	userID, err := h.service.CreateUser(user)
	if err != nil {
		logger.Logger.Error("error creating user", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, newErrorResponse("error creating user"))
		return
	}

	logger.Logger.Info("user created")
	logger.Logger.Debug(fmt.Sprintf("user with id %d created", userID))
	c.JSON(http.StatusCreated, newSuccessResponse(strconv.Itoa(userID)))
}

// UpdateUser updates the user
// @Summary Update a user
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param request body model.UserUpdateRequestBody true "User details"
// @Success 200 {object} SuccessResponse "Message"
// @Failure 400 {object} ErrorResponse "Error message"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /user/{user_id} [patch]
func (h *userHandler) UpdateUser(c *gin.Context) {
	logger.Logger.Info("start update user")
	id := c.Param("user_id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		logger.Logger.Error("error parsing user id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("error parsing user id"))
		return
	}
	var input model.UserUpdateRequestBody
	if err := c.BindJSON(&input); err != nil {
		logger.Logger.Error("error binding json", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("error binding json"))
		return
	}
	logger.Logger.Debug("parsed input", slog.Any("input", input))

	user, err := h.service.GetUser(userID)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			logger.Logger.Warn("user not found", slog.Int("user_id", userID))
			c.JSON(http.StatusBadRequest, newErrorResponse("user not found"))
			return
		}
		logger.Logger.Error("error getting user", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, newErrorResponse("error getting user"))
		return
	}
	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Surname != nil {
		user.Surname = *input.Surname
	}
	if input.Patronymic != nil {
		user.Patronymic = *input.Patronymic
	}
	if input.Address != nil {
		user.Address = *input.Address
	}

	if err := h.service.UpdateUser(user); err != nil {
		logger.Logger.Error("error updating user", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, newErrorResponse("error updating user"))
		return
	}

	logger.Logger.Info("user updated")
	logger.Logger.Debug(fmt.Sprintf("user with id %d updated", userID))
	c.JSON(200, newSuccessResponse("user updated"))
}

// DeleteUser deletes the user
// @Summary Delete a user
// @Tags Users
// @Param user_id path int true "User ID"
// @Success 200 {object} SuccessResponse "Message"
// @Failure 400 {object} ErrorResponse "Error message"
// @Failure 500 {object} ErrorResponse "Error message"
// @Router /user/{user_id} [delete]
func (h *userHandler) DeleteUser(c *gin.Context) {
	logger.Logger.Info("start delete user")

	id := c.Param("user_id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		logger.Logger.Error("error parsing user id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("error parsing user id"))
		return
	}

	if err := h.service.DeleteUser(userID); err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			logger.Logger.Warn("user not found", slog.Int("user_id", userID))
			c.JSON(http.StatusBadRequest, newErrorResponse("user not found"))
			return
		}
		logger.Logger.Error("error deleting user", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, newErrorResponse("error deleting user"))
		return
	}

	logger.Logger.Info("user deleted")
	logger.Logger.Debug(fmt.Sprintf("user with id %d deleted", userID))
	c.JSON(http.StatusOK, newSuccessResponse("user deleted"))
}

func parsePassportNumberAndSerie(passportData string) (int, int, error) {
	splitNumber := strings.Split(passportData, " ")
	if len(splitNumber) != 2 {
		return 0, 0, errors.New("invalid passport number")
	}

	passportSerie, err := strconv.Atoi(splitNumber[0])
	if err != nil {
		return 0, 0, errors.New("invalid passport serie")
	}

	passportNumber, err := strconv.Atoi(splitNumber[1])
	if err != nil {
		return 0, 0, errors.New("invalid passport number")
	}

	return passportSerie, passportNumber, nil
}
