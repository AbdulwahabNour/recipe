package http

import (
	"github.com/AbdulwahabNour/recipe/internal/recipe"
	"github.com/gin-gonic/gin"
)

func RecipeRoutes(r *gin.RouterGroup, handler recipe.Handlers) {
	r.POST("/recipes", handler.CreateRecipeHandler)
	r.GET("/recipes", handler.ListRecipeHandler)
	r.PUT("/recipes", handler.UpdateRecipeHandler)
	r.DELETE("/recipes/:id", handler.DeleteRecipeHandler)
}
