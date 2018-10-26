package main
import (
 "github.com/labstack/echo"
 "github.com/labstack/echo/engine/standard"
 "github.com/labstack/echo/middleware"
)
func main() {
 e := echo.New()
// Middleware
 e.Use(middleware.Logger())
 e.Use(middleware.Recover())
//CORS
 e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  AllowOrigins: []string{"*"},
  AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
 }))
// Route => handler  e.GET("/", func(c echo.Context) error {       return c.String(http.StatusOK, "Hello, World!\n")  
})
// Server
 e.Run(standard.New(":1323"))
}
// Route => handler  e.GET("/", func(c echo.Context) error {       return c.String(http.StatusOK, "Hello, World!\n")  
})