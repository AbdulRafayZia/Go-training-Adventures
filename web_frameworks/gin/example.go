// package main

// import (
//   "net/http"

//   "github.com/gin-gonic/gin"
// )

// func main() {
//   r := gin.Default()
//   r.GET("/ping", func(c *gin.Context) {
//     c.JSON(http.StatusOK, gin.H{
//       "message": "pong",
//     })
//   })
//   r.Run( ) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
    // Create a new Gin router
    router := gin.Default()

    // Define a route and handler for the "/" path
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello world!",
        })
    })

    // Start the server on port 8080
    router.Run(":8080")
}
