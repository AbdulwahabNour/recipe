package recipe

import "github.com/gin-gonic/gin"

type Handlers interface {
	CreateRecipeHandler(c *gin.Context)
	ListRecipeHandler(c *gin.Context)
	UpdateRecipeHandler(c *gin.Context)
	DeleteRecipeHandler(c *gin.Context)
	SearchRecipeHandler(c *gin.Context)
}
