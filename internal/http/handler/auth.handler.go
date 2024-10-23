package handler

import (
	"net/http"

	"github.com/Sandhya-Pratama/Ngonser-App/entity"
	"github.com/Sandhya-Pratama/Ngonser-App/internal/http/validator"
	"github.com/Sandhya-Pratama/Ngonser-App/internal/service"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	registrationService service.RegistrationUseCase
	loginService        service.LoginUseCase
	tokenService        service.TokenUseCase
}

func NewAuthHandler(
	registrationService service.RegistrationUseCase,
	loginService service.LoginUseCase,
	tokenService service.TokenUseCase,
) *AuthHandler {

	return &AuthHandler{
		registrationService: registrationService,
		loginService:        loginService,
		tokenService:        tokenService,
	}
}

func (h *AuthHandler) Login(ctx echo.Context) error {
	//pengecekan request
	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	// di cek pake validate buat masukin input
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	//untuk manggil login service di folder service
	user, err := h.loginService.Login(ctx.Request().Context(), input.Email, input.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	//untuk manggil token service di folder service
	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	data := map[string]interface{}{
		"access_token": accessToken,
	}
	return ctx.JSON(http.StatusOK, data)
}

// Public Register
func (h *AuthHandler) Registration(ctx echo.Context) error {
	//pengecekan request
	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Roles    string `json:"roles" validate:"required"`
		Number   string `json:"number" validate:"required,min=11,max=13"`
	}

	if err := ctx.Bind(&input); err != nil { // di cek pake validate buat masukin input
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	//untuk manggil registration service di folder service
	user := entity.Register(input.Email, input.Password, input.Roles, input.Number)
	err := h.registrationService.Registration(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":      "User registration successfully",
		"access_token": accessToken,
	})

}
