package router

import (
	"my-favorite-pokemon-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, sc controller.IStarController) *echo.Echo {
	e := echo.New()
	// レートリミットに引っかからない。とりあえず、後々対応する。（30分で10回はテスト用）
	// e.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
	// 	Skipper: middleware.DefaultSkipper,
	// 	Store: middleware.NewRateLimiterMemoryStoreWithConfig(
	// 		middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 10, ExpiresIn: 30 * time.Minute},
	// 	),
	// 	IdentifierExtractor: func(ctx echo.Context) (string, error) {
	// 		id := ctx.RealIP()
	// 		return id, nil
	// 	},
	// }))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// CORS設定
		AllowOrigins: []string{"https://localhost:3000", os.Getenv("FE_URL"), "https://my-favorite-pokemon-app-front.vercel.app"}, //Vercel,Renderの環境ではCORSに引っかかるため無理やり設定
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		// デフォルト
		// DefaultCSRFConfig = CSRFConfig{
		// 	Skipper:      DefaultSkipper,
		// 	TokenLength:  32,
		// 	TokenLookup:  "header:" + echo.HeaderXCSRFToken,
		// 	ContextKey:   "csrf",
		// 	CookieName:   "_csrf",
		// 	CookieMaxAge: 86400, //秒
		// }
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode, //PostMan確認用。（SameSiteNoneModeだとSecureが自動でtrueになるため）
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	s := e.Group("/stars")
	s.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		// TokenLookup: Default value "header:Authorization"
		TokenLookup: "cookie:token",
	}))
	s.GET("", sc.GetAllStars)
	s.POST("", sc.CreateStar)
	s.PUT("/:pokemonId", sc.UpdateStar)
	s.DELETE("/:pokemonId", sc.DeleteStar)
	e.Use(middleware.Logger())
	return e
}
