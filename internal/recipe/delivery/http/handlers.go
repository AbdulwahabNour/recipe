package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	model "github.com/AbdulwahabNour/recipe/internal/model/recipe"
	"github.com/AbdulwahabNour/recipe/internal/recipe"
	"github.com/gin-gonic/gin"
)

type apiHandler struct {
	service recipe.Service
}

func NewRecipeHandler(serv recipe.Service) recipe.Handlers {
	return &apiHandler{
		service: serv,
	}
}

// @Summary      Create a New Recipe
// @Description  Create a new recipe by providing Recipe JSON data and store it in the database. Returns the saved recipe as JSON.
// @Produce      json
// @Param        Recipe  body  model.RecipeForm  true  "Recipe JSON"
// @Success      200   {object}   model.Recipe
// @failure      400   {object}  map[string]string
// @failure      500   {object}  map[string]string
// @Router       /recipes [post]
func (h *apiHandler) CreateRecipeHandler(c *gin.Context) {
	var recForm model.RecipeForm
	if err := c.ShouldBindJSON(&recForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()
	recipe := recForm.ToRecipe()
	err := h.service.CreateRecipe(ctx, recipe)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting a new recipe"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

// @Summary      List all Recipes
// @Description  List recipes.
// @Produce      json
// @Success      200  {object}  []model.Recipe
// @failure      500   {object}  map[string]string
// @Router       /recipes [get]
func (h *apiHandler) ListRecipeHandler(c *gin.Context) {
	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()
	recipes, err := h.service.ListAllRecipes(ctx)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error happened while retrieving recipes"})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

// @Summary      Update Recipe
// @Description  Update an existing recipe by providing updated Recipe JSON data. Returns updated recipe upon successful update.
// @Produce      json
// @Param        Recipe  body  model.Recipe  true  "Updated Recipe JSON"
// @Success      200   {object}  map[string]string
// @failure      400   {object}  map[string]string
// @failure      500   {object}  map[string]string
// @Router       /recipes [put]
func (h *apiHandler) UpdateRecipeHandler(c *gin.Context) {

	var recipe model.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()
	err := h.service.UpdateRecipe(ctx, &recipe)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error happened while updating recipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe has been updated"})
}

// @Summary      Delete Recipe
// @Description  Delete a recipe by providing its unique identifier (ID). Returns a confirmation message upon successful deletion.
// @Produce      json
// @Param        id  path  string  true  "Recipe ID"
// @Success      200   {object}   map[string]string
// @failure      500   {object}  map[string]string
// @Router       /recipes/{id} [delete]
func (h *apiHandler) DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()
	err := h.service.DeleteRecipe(ctx, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error happened while deleting recipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Recipe with id (%s) has been deleted", id)})
}

// @Summary      Search Recipes
// @Description  Search for recipes based on specified tags. Returns a list of matching recipes.
// @Produce      json
// @Param        tags  query  string  true  "Comma-separated list of tags for filtering"
// @Success      200   {array}   model.Recipe
// @failure      500   {object}  map[string]string
// @Router       /recipes/search [get]
func (h *apiHandler) SearchRecipeHandler(c *gin.Context) {
	tags := strings.Split(c.Query("tags"), ",")
	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()
	recipes, err := h.service.SearchRecipe(ctx, tags)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error happened while deleting recipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recipes": recipes})
}
