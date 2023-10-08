package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AbdulwahabNour/recipe/docs"
	recipeHandlers "github.com/AbdulwahabNour/recipe/internal/recipe/delivery/http"
	recipeRepository "github.com/AbdulwahabNour/recipe/internal/recipe/repo/mongo"
	recipeService "github.com/AbdulwahabNour/recipe/internal/recipe/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	maxHeaderBytes = 1 << 20 //1MB
)

type server struct {
	ginEngin *gin.Engine
	mongoDB  *mongo.Database
}

func NewServer(mongodb *mongo.Database) *server {
	return &server{
		ginEngin: gin.New(),
		mongoDB:  mongodb,
	}
}

func (s *server) MapHandler() error {
	docs.SwaggerInfo.BasePath = "/api/v1"
	s.ginEngin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	dbCollection := s.mongoDB.Collection("recipe")

	recipeRepo := recipeRepository.NewRecipeRepo(dbCollection)
	recipeServ := recipeService.NewRecipeService(recipeRepo)
	recipeHandler := recipeHandlers.NewRecipeHandler(recipeServ)

	v1 := s.ginEngin.Group("/api/v1")
	recipeHandlers.RecipeRoutes(v1, recipeHandler)

	return nil
}

func (s *server) Run() error {
	err := s.MapHandler()

	if err != nil {
		return err
	}
	srv := &http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        s.ginEngin,
		MaxHeaderBytes: maxHeaderBytes,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	msg := <-c
	log.Printf("Server exiting with signal %s", msg)
	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()

	err = srv.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("Server exiting")
	return nil
}
