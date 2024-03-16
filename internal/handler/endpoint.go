package handler

import (
	"net/http"
	"strconv"

	"github.com/luthfirahman/user-svc/internal/dto"
	"github.com/luthfirahman/user-svc/internal/service"
	"github.com/luthfirahman/user-svc/internal/utils"
	"github.com/luthfirahman/user-svc/internal/utils/constant"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetMyProfile(ctx *gin.Context)
	UpdateMyProfile(ctx *gin.Context)
}

type handler struct {
	userService service.IUserService
}

func NewHandler(userService service.IUserService) Handler {
	return &handler{
		userService: userService,
	}
}

func (h handler) Register(ctx *gin.Context) {
	var r *dto.RegisterRequest

	if err := ctx.ShouldBind(&r); err != nil {
		res := utils.BuildResponse(constant.MESSAGE_FAILED_GET_DATA, nil, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	validationErr := r.Validate()
	if (validationErr != dto.RegisterValidationErrorResponse{}) {
		res := utils.BuildResponse(constant.MESSAGE_INCORRECT_DATA, nil, validationErr)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := h.userService.Register(ctx.Request.Context(), r)
	if err != nil {
		res := utils.BuildResponse(err.Error(), nil, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data := dto.RegisterResponse{ID: result.Id}
	res := utils.BuildResponse(constant.MESSAGE_SUCCESS_REGISTER_USER, data, nil)
	ctx.JSON(http.StatusOK, res)
}

func (h handler) Login(ctx *gin.Context) {
	var r dto.LoginRequest

	if err := ctx.ShouldBind(&r); err != nil {
		res := utils.BuildResponse(constant.MESSAGE_FAILED_GET_DATA, nil, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := h.userService.Login(ctx.Request.Context(), r)
	if err != nil {
		res := utils.BuildResponse(err.Error(), nil, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data := dto.LoginResponse{ID: result.ID, Token: result.Token}
	res := utils.BuildResponse(constant.MESSAGE_SUCCESS_LOGIN, data, nil)
	ctx.JSON(http.StatusOK, res)
}

func (h handler) GetMyProfile(ctx *gin.Context) {

	userIDStr := ctx.MustGet("user_id").(string)
	userID, _ := strconv.Atoi(userIDStr)

	result, err := h.userService.GetMyProfile(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponse(err.Error(), nil, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data := dto.ProfileResponse{
		Name:        result.Name,
		PhoneNumber: result.PhoneNumber,
	}

	res := utils.BuildResponse(constant.MESSAGE_SUCCESS, data, nil)
	ctx.JSON(http.StatusOK, res)
}

func (h handler) UpdateMyProfile(ctx *gin.Context) {
	var r *dto.ProfileUpdateRequest

	userIDStr := ctx.MustGet("user_id").(string)
	userID, _ := strconv.Atoi(userIDStr)

	if err := ctx.ShouldBind(&r); err != nil {
		res := utils.BuildResponse(constant.MESSAGE_FAILED_GET_DATA, nil, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	validationErr := r.Validate()
	if (validationErr != dto.RegisterValidationErrorResponse{}) {
		res := utils.BuildResponse(constant.MESSAGE_INCORRECT_DATA, nil, validationErr)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := h.userService.UpdateMyProfile(ctx.Request.Context(), r, userID)
	if err != nil {
		res := utils.BuildResponse(err.Error(), nil, nil)
		if err.Error() == "phone number not available" {
			ctx.JSON(http.StatusConflict, res)
		} else {
			ctx.JSON(http.StatusBadRequest, res)

		}
		return
	}

	data := dto.ProfileResponse{
		Name:        result.Name,
		PhoneNumber: result.PhoneNumber,
	}

	res := utils.BuildResponse(constant.MESSAGE_SUCCESS, data, nil)
	ctx.JSON(http.StatusOK, res)
}
