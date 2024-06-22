package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Application struct {
	logger echo.Logger // จัดการข้อความ log
	server *echo.Echo  //จัดการ routing , middleware , request
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5173", "http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Public routes

	e.POST("/loginuser", loginUserHandler)
	e.POST("/registerclient", saveUserHandler)

	userGroup := e.Group("/userdashboard")
	userGroup.Use(jwtMiddleware, userMiddleware)
	userGroup.GET("/:idnumber", fetchUserHandler)

	adminGroup := e.Group("/admindashboard")
	adminGroup.Use(jwtMiddleware, adminMiddleware) //handler function จากหน้าไปหลัง
	adminGroup.GET("", fetchAdminHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Protected routes
	e.GET("/userinfo", getUserInfoHandler, jwtMiddleware) //handler function จากหลังไปหน้า

	// Assigning logger and server to Application struct
	app := Application{
		logger: e.Logger,
		server: e,
	}

	fmt.Println(app)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
