package mongorepo

import (
	"context"
	"fmt"
	"time"

	model "github.com/AbdulwahabNour/recipe/internal/model/recipe"
	"github.com/AbdulwahabNour/recipe/internal/recipe"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type recipeRepo struct {
	collection *mongo.Collection
}

func NewRecipeRepo(c *mongo.Collection) recipe.Repository {
	return &recipeRepo{
		collection: c,
	}
}

func (r *recipeRepo) CreateRecipe(ctx context.Context, recipe *model.Recipe) error {
	recipe.ID = primitive.NewObjectID()
	recipe.PublishedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, recipe)
	if err != nil {
		return err
	}
	return nil
}
func (r *recipeRepo) ListAllRecipes(ctx context.Context) ([]*model.Recipe, error) {
	recipes := make([]*model.Recipe, 0)
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var recipe *model.Recipe
		cur.Decode(recipe)
		recipes = append(recipes, recipe)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}
func (r *recipeRepo) UpdateRecipe(ctx context.Context, recipe *model.Recipe) error {

	// collection.UpdateOne( {_id: id of this doc} ,{$set:{ name:"new name", instu}})
	filter := bson.M{"_id": recipe.ID}
	recipeUpdate := bson.D{
		{"$set", bson.D{
			{"name", recipe.Name},
			{"instructions", recipe.Instructions},
			{"ingredients", recipe.Ingredients},
			{"tags", recipe.Tags}},
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, recipeUpdate)

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {

		return fmt.Errorf("No documents were updated.")
	}

	return nil
}
func (r *recipeRepo) DeleteRecipe(ctx context.Context, id string) error {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": idObj}
	result, err := r.collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {

		return fmt.Errorf("No documents were deleted.")
	}

	return nil
}
func (r *recipeRepo) SearchRecipe(ctx context.Context, tags []string) ([]*model.Recipe, error) {
	filter := bson.M{"tags": bson.M{"$all": tags}}
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	recipes := make([]*model.Recipe, 0)
	for cur.Next(ctx) {
		var recipe *model.Recipe
		cur.Decode(recipe)
		recipes = append(recipes, recipe)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}
