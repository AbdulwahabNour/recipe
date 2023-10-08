package service

import (
	"context"

	model "github.com/AbdulwahabNour/recipe/internal/model/recipe"
	"github.com/AbdulwahabNour/recipe/internal/recipe"
)

type recipeService struct {
	repo recipe.Repository
}

func NewRecipeService(repo recipe.Repository) recipe.Service {
	return &recipeService{
		repo: repo,
	}
}

func (s *recipeService) CreateRecipe(ctx context.Context, recipe *model.Recipe) error {

	err := s.repo.CreateRecipe(ctx, recipe)

	if err != nil {
		return err
	}

	return nil
}
func (s *recipeService) ListAllRecipes(ctx context.Context) ([]*model.Recipe, error) {

	return s.repo.ListAllRecipes(ctx)
}
func (s *recipeService) UpdateRecipe(ctx context.Context, recipe *model.Recipe) error {

	err := s.repo.UpdateRecipe(ctx, recipe)
	if err != nil {
		return err
	}

	return nil
}
func (s *recipeService) DeleteRecipe(ctx context.Context, id string) error {
	err := s.repo.DeleteRecipe(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
func (s *recipeService) SearchRecipe(ctx context.Context, tags []string) ([]*model.Recipe, error) {

	return s.repo.SearchRecipe(ctx, tags)
}
