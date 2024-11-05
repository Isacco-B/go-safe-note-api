package routes

import (
    "github.com/Isacco-B/go-safe-note-api/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    userGroup := r.Group("/notes")
    {
        userGroup.POST("/", controllers.CreateNote)
        userGroup.GET("/:link", controllers.GetNote)
    }

    return r
}
