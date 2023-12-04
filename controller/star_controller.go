package controller

import (
	"my-favorite-pokemon-rest-api/model"
	"my-favorite-pokemon-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Interface
type IStarController interface {
	GetAllStars(c echo.Context) error
	CreateStar(c echo.Context) error
	UpdateStar(c echo.Context) error
	DeleteStar(c echo.Context) error
}

// interfaceを実装するstruct
type starController struct {
	su usecase.IStarUsecase
}

// コンストラクト
func NewStarController(su usecase.IStarUsecase) IStarController {
	return &starController{su: su}
}

// 実装部
func getUserIDFromJWT(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["user_id"].(float64)

	return uint(user_id)
}

func (sc *starController) GetAllStars(c echo.Context) error {
	user_id := getUserIDFromJWT(c)
	stars, err := sc.su.GetAllStars(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, stars)
}
func (sc *starController) CreateStar(c echo.Context) error {
	user_id := getUserIDFromJWT(c)
	star := model.Star{}
	if err := c.Bind(&star); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	star.UserID = user_id //後で確認
	stars, err := sc.su.CreateStar(star)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, stars)
}
func (sc *starController) UpdateStar(c echo.Context) error {
	user_id := getUserIDFromJWT(c)
	s_pokemon_id := c.Param("pokemonId")
	pokemon_id, err := strconv.Atoi(s_pokemon_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	star := model.Star{}
	if err := c.Bind(&star); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	star.UserID = user_id //後で確認
	resStar, err := sc.su.UpdateStar(star, uint(pokemon_id), user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resStar)
}
func (sc *starController) DeleteStar(c echo.Context) error {
	user_id := getUserIDFromJWT(c)
	s_pokemon_id := c.Param("pokemonId")
	pokemon_id, err := strconv.Atoi(s_pokemon_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := sc.su.DeleteStar(uint(pokemon_id), user_id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
