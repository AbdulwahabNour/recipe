package http

import (
	"fmt"
	"net/http"
	"time"

	model "github.com/AbdulwahabNour/recipe/internal/model/recipe"
	"github.com/AbdulwahabNour/recipe/internal/recipe"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type apiHandler struct {
}

func NewRecipeHandler() recipe.Handlers {
	return &apiHandler{}
}
func (h *apiHandler) CreateRecipeHandler(c *gin.Context) {
	var rec model.Recipe
	if err := c.ShouldBindJSON(&rec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rec.ID = xid.New().String()
	rec.PublishedAt = time.Now()
	//sent to service

	c.JSON(http.StatusOK, rec)
}
func (h *apiHandler) ListRecipeHandler(c *gin.Context) {

	c.JSON(http.StatusOK, []model.Recipe{})
}
func (h *apiHandler) UpdateRecipeHandler(c *gin.Context) {
	var recipe model.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//sent to service
	c.JSON(http.StatusOK, []model.Recipe{})
}
func (h *apiHandler) DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")

	//sent to service

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Recipe with id (%s) has been deleted", id)})
}
func (h *apiHandler) SearchRecipeHandler(c *gin.Context) {
	id := c.Query("tags")

	//sent to service

	c.JSON(http.StatusOK, gin.H{"message": []model.Recipe{{ID: id}}})
}
