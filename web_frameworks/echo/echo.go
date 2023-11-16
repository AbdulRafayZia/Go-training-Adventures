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
    e.GET("/msg/:msg", func (c echo.Context) error {
        // User ID from path `users/:id`
        id := c.Param("msg")
        return c.String(http.StatusOK,"msg  :"+ id)
    })
    // e.GET("/users/:id", getUser)

    e.Logger.Fatal(e.Start(":1323"))
  

    // $ curl -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:1323/save
// => name:Joe Smith, email:joe@labstack.com
 
}