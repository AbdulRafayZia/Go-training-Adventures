package main

import (
    "net/http"
    
    "github.com/labstack/echo/v4"
)

type User struct {
    Name  string `json:"name" xml:"name" form:"name" query:"name"`
    Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
    e.POST("/users", func(c echo.Context) error {
        u := new(User)
        if err := c.Bind(u); err != nil {
            return err
        }
        return c.JSON(http.StatusCreated, u)
        // or
        // return c.XML(http.StatusCreated, u)
    })
}