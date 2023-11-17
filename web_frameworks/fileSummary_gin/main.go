package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
	"main.go/pkg"
)


type Arguments struct {
    FilePath     string  `json:"id"`
    Routines  int  `json:"title"`
}

func main() {

	router := gin.Default()
 
    router.POST("/albums", PostArguments)

    router.Run("localhost:8080")


}

func PostArguments(c *gin.Context) {
    var newArgs  Arguments

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newArgs); err != nil {
        return
    }

    // Add the new album to the slice.
    // albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, pkg.ProcessFile( newArgs.FilePath , newArgs.Routines) )
}


