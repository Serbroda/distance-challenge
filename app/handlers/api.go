package handlers

import (
	"errors"
	"github.com/Serbroda/distance-challenge/models"
	"github.com/Serbroda/distance-challenge/security"
	"github.com/Serbroda/distance-challenge/services"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	ErrParamNotFound = errors.New("param not found")
)

type UserChange struct {
	Password string `json:"password" form:"password"`
}

type UserHandler struct {
	UserService *services.UserService
	RunService  *services.RunService
}

func RegisterUserHandlers(e *echo.Group, h *UserHandler, baseUrl string) {
	e.GET(baseUrl+"/me", h.Me)
	e.GET(baseUrl+"/users", h.GetUsers)
	e.GET(baseUrl+"/users/:id", h.GetUser)
	e.PUT(baseUrl+"/users/:id", h.UpdateUser)
	e.POST(baseUrl+"/users/:id/activate", h.Activate)
	e.POST(baseUrl+"/users/:id/runs", h.CreateRun)
	e.GET(baseUrl+"/runs", h.GetRuns)
	e.GET(baseUrl+"/runs/:id", h.GetRun)
	e.PUT(baseUrl+"/users/:runId", h.UpdateRun)
	e.DELETE(baseUrl+"/users/:runId", h.DeleteRun)
}

// @Tags api
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /me [get]
func (h *UserHandler) Me(ctx echo.Context) error {
	u, err := h.getUser(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, u)
}

// @Tags api
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /users [get]
func (h *UserHandler) GetUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.UserService.GetAll())
}

// @Tags api
// @Accept json
// @Produce json
// @Param userId path string false "user id"
// @Success 200 {object} models.User
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /users/{userId} [get]
func (h *UserHandler) GetUser(ctx echo.Context) error {
	userId64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return echo.ErrInternalServerError
	}
	userId := uint(userId64)

	user, err := h.UserService.Get(userId)
	if err != nil {
		return echo.ErrNotFound
	}
	return ctx.JSON(http.StatusOK, user)
}

// @Tags api
// @Accept json
// @Produce json
// @Param userId path string false "user id"
// @Param user body models.User true "user body"
// @Success 200 {object} models.User
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /users/{userId} [put]
func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	u, err := h.getUser(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}

	userId, err := h.parseParamUint(ctx, "id")
	if err != nil {
		return err
	}

	var payload UserChange
	err = ctx.Bind(&payload)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	if u.ID != userId {
		return echo.ErrForbidden
	}

	user, err := h.UserService.Get(userId)
	if err != nil {
		return echo.ErrNotFound
	}
	user.Password = security.MustHashPassword(payload.Password)
	user, err = h.UserService.Update(user)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, user)
}

// @Tags api
// @Accept json
// @Produce json
// @Param userId path string false "user id"
// @Param user body models.User true "user body"
// @Success 200 {object} models.User
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /users/{userId}/activate [put]
func (h *UserHandler) Activate(ctx echo.Context) error {
	u, err := h.getUser(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}
	if !u.IsAdmin {
		return echo.NewHTTPError(http.StatusForbidden, "no admin")
	}
	userId := ctx.Param("id")
	return ctx.String(http.StatusOK, "user "+userId+" activated")
}

// @Tags api
// @Accept json
// @Produce json
// @Success 200 {array} models.Run
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /runs [get]
func (h *UserHandler) GetRuns(ctx echo.Context) error {
	var runs []models.Run
	userId, err := h.parseParamUint(ctx, "userId")
	if err != nil {
		if err == ErrParamNotFound {
			runs = h.RunService.GetAll()
		} else {
			return echo.ErrInternalServerError
		}
	}
	runs, err = h.RunService.GetByUser(userId)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, runs)
}

// @Tags api
// @Accept json
// @Produce json
// @Param id path string false "user id"
// @Param user body models.Run true "run body"
// @Success 200 {array} models.Run
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /users/{id}/runs [post]
func (h *UserHandler) CreateRun(ctx echo.Context) error {
	userId, err := h.parseParamUint(ctx, "id")
	if err != nil {
		return err
	}
	var payload models.Run
	err = ctx.Bind(&payload)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	payload.UserId = userId

	run, err := h.RunService.Create(payload)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, run)
}

// @Tags api
// @Accept json
// @Produce json
// @Param id path string false "run id"
// @Success 200 {object} models.Run
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /runs/{id} [get]
func (h *UserHandler) GetRun(ctx echo.Context) error {
	rundId, err := h.parseParamUint(ctx, "runId")
	if err != nil {
		return err
	}

	run, err := h.RunService.Get(rundId)
	if err != nil {
		return echo.ErrNotFound
	}
	return ctx.JSON(http.StatusOK, run)
}

// @Tags api
// @Accept json
// @Produce json
// @Param id path string false "run id"
// @Param user body models.Run true "run body"
// @Success 200 {object} models.Run
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /runs/{id} [put]
func (h *UserHandler) UpdateRun(ctx echo.Context) error {
	rundId, err := h.parseParamUint(ctx, "runId")
	if err != nil {
		return echo.ErrInternalServerError
	}
	u, err := h.getUser(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}
	run, err := h.RunService.Get(rundId)
	if err != nil {
		return echo.ErrInternalServerError
	}
	if u.ID != run.UserId {
		return echo.ErrForbidden
	}

	var payload models.Run
	err = ctx.Bind(&payload)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	run.Distance = payload.Distance
	run.Time = payload.Time
	run, err = h.RunService.Update(run)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, run)
}

// @Tags api
// @Accept json
// @Produce json
// @Param id path string false "run id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /runs/{id} [delete]
func (h *UserHandler) DeleteRun(ctx echo.Context) error {
	rundId, err := h.parseParamUint(ctx, "runId")
	if err != nil {
		return echo.ErrInternalServerError
	}
	u, err := h.getUser(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}
	run, err := h.RunService.Get(rundId)
	if err != nil {
		return echo.ErrInternalServerError
	}
	if u.ID != run.UserId {
		return echo.ErrForbidden
	}

	err = h.RunService.DeleteRun(rundId)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.String(http.StatusOK, "successfully deleted")
}

func (si *UserHandler) getUser(ctx echo.Context) (models.User, error) {
	obj := ctx.Get("user")
	token := obj.(*jwt.Token)
	claims := token.Claims.(*security.Claims)
	if claims == nil {
		return models.User{}, echo.ErrUnauthorized
	}
	return si.UserService.Get(claims.UserId)
}

func (si *UserHandler) parseParamUint(ctx echo.Context, param string) (uint, error) {
	p := ctx.Param(param)
	if p == "" {
		return 0, ErrParamNotFound
	}
	userId64, err := strconv.ParseUint(p, 10, 64)
	if err != nil {
		return 0, echo.ErrInternalServerError
	}
	return uint(userId64), nil
}
