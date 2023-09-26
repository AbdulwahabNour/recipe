package recipe

import "github.com/gin-gonic/gin"

type Handlers interface {
	CreateRecipeHandler(c *gin.Context)
}
