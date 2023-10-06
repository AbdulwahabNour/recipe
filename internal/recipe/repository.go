package recipe

import (
	"context"

	model "github.com/AbdulwahabNour/recipe/internal/model/recipe"
)

type Repository interface {
	CreateRecipe(ctx context.Context, recipe *model.Recipe) error
	ListAllRecipes(ctx context.Context) ([]*model.Recipe, error)
	UpdateRecipe(ctx context.Context, recipe *model.Recipe) error
	DeleteRecipe(ctx context.Context, id string) error
	SearchRecipe(ctx context.Context, tags []string) ([]*model.Recipe, error)
}
