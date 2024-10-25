package infra

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5"

	"handler/function/domain/router"
	"handler/function/domain/service"
	"handler/function/infra/assets"
	"handler/function/infra/config"
	"handler/function/infra/mongo"
	"handler/function/infra/repository"
)

func Init() (*chi.Mux, func()) {
	cfg := config.GetConfig()

	mongoClient := mongo.Init()

	repo := repository.NewRepository(mongoClient.Database(cfg.MongoDB))
	s := service.NewService(repo)

	r := chi.NewRouter()
	router := router.NewRouter(s)
	router.Mount(r)
	assets.Mount(r)

	return r, func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}
}
