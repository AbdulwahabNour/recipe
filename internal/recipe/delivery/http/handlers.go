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

// @Summary      Create a New Recipe
// @Description  Create a new recipe by providing Recipe JSON data and store it in the database. Returns the saved recipe as JSON.
// @Produce      json
// @Param        Recipe  body  model.Recipe  true  "Recipe JSON"
// @Success      200   {object}   model.Recipe
// @Router       /recipes [post]
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

// @Summary      Update Recipe
// @Description  Update an existing recipe by providing updated Recipe JSON data. Returns updated recipe upon successful update.
// @Produce      json
// @Param        Recipe  body  model.Recipe  true  "Updated Recipe JSON"
// @Success      200  {object}  model.Recipe
// @Router       /recipes [put]
func (h *apiHandler) UpdateRecipeHandler(c *gin.Context) {

	var recipe model.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//sent to service
	c.JSON(http.StatusOK, model.Recipe{})
}

// @Summary      Delete Recipe
// @Description  Delete a recipe by providing its unique identifier (ID). Returns a confirmation message upon successful deletion.
// @Produce      json
// @Param        id  path  string  true  "Recipe ID"
// @Success      200   {object}   map[string]interface{}
// @Router       /recipes/{id} [delete]
func (h *apiHandler) DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")

	//sent to service

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Recipe with id (%s) has been deleted", id)})
}

// @Summary      Search Recipes
// @Description  Search for recipes based on specified tags. Returns a list of matching recipes.
// @Produce      json
// @Param        tags  query  string  true  "Comma-separated list of tags for filtering"
// @Success      200   {array}   model.Recipe
// @Router       /recipes/search [get]
func (h *apiHandler) SearchRecipeHandler(c *gin.Context) {
	id := c.Query("tags")

	//sent to service

	c.JSON(http.StatusOK, gin.H{"message": []model.Recipe{{ID: id}}})
}
