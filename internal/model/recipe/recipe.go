package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" `
	Tags         []string           `json:"tags" `
	Ingredients  []string           `json:"ingredients" `
	Instructions []string           `json:"instructions" `
	PublishedAt  time.Time          `json:"publishedAt" `
}

type RecipeForm struct {
	Name         string   `json:"name" `
	Tags         []string `json:"tags" `
	Ingredients  []string `json:"ingredients" `
	Instructions []string `json:"instructions" `
}

func (r RecipeForm) ToRecipe() *Recipe {
	return &Recipe{
		Name:         r.Name,
		Tags:         r.Tags,
		Ingredients:  r.Ingredients,
		Instructions: r.Instructions,
	}
}
