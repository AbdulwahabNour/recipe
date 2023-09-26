package http

import (
	"github.com/AbdulwahabNour/recipe/internal/recipe"
	"github.com/gin-gonic/gin"
)

func RecipeRoutes(r *gin.RouterGroup, handler recipe.Handlers) {
	r.POST("/recipes", handler.CreateRecipeHandler)
}