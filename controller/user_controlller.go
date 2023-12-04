package controller

import (
	"my-favorite-pokemon-rest-api/model"
	"my-favorite-pokemon-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// interface
type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

// interfaceを実装するstruct
type userController struct {
	uu usecase.IUserUsecase
}

// コンストラクタ
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu: uu}
}

// メソッドの処理内容
func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	resUser, err := uc.uu.SingUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resUser)
}

func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	token, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		Expires:  time.Now().Add(time.Hour * 12),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode, // フロントがSPAのため、Noneとする。
	}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, token)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := &http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/",
		Domain:  os.Getenv("API_DOMAIN"),
		Expires: time.Now(),
		// Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode, // フロントがSPAのため、Noneとする。
	}
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
