package router

import (
	"encoding/gob"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/jezpoz/device-camera-web/platform/authenticator"
	"github.com/jezpoz/device-camera-web/platform/middleware"
	"github.com/jezpoz/device-camera-web/web/app/callback"
	"github.com/jezpoz/device-camera-web/web/app/home"
	"github.com/jezpoz/device-camera-web/web/app/login"
	"github.com/jezpoz/device-camera-web/web/app/logout"
	"github.com/jezpoz/device-camera-web/web/app/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug"
	}
	gin.SetMode(ginMode)

	router := gin.New()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", home.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
