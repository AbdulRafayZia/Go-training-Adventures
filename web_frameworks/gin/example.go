package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
    c.JSON(http.StatusOK, albums)
}
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
// package main

// import (
//   "net/http"

//   "github.com/gin-gonic/gin"
// )

// func main() {
//     // Create a new Gin router
//     router := gin.Default()

//     // Define a route and handler for the "/" path
//     router.GET("/", func(c *gin.Context) {
//         c.JSON(http.StatusOK, gin.H{
//             "message": "Hello world!",
//         })
//     })

//     // Start the server on port 8080
//     router.Run(":8080")
// }
