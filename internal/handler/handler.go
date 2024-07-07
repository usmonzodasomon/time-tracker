package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/usmonzodasomon/time-tracker/docs"
)

func NewRouter(handler *gin.Engine, db *sqlx.DB) {
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	h := handler.Group("/api")
	{
		h.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		newUserHandler(h, db)
		newTaskHandler(h, db)

	}
}
