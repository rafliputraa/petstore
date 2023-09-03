// Package app configures and runs application.
package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rafliputraa/petstore/config"
	v1 "github.com/rafliputraa/petstore/internal/controller/v1"
	"github.com/rafliputraa/petstore/internal/repo"
	usecase_customer "github.com/rafliputraa/petstore/internal/usecase/customer"
	usecase_pet "github.com/rafliputraa/petstore/internal/usecase/pet"
	"github.com/rafliputraa/petstore/pkg/logger"
	"github.com/rafliputraa/petstore/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) *gin.Engine {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New("postgres://"+cfg.PG.Username+":"+cfg.PG.Password+"@"+cfg.PG.Host+":"+cfg.PG.Port+"/"+cfg.PG.DbName, postgres.MaxPoolSize(cfg.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	customerUseCase := usecase_customer.New(
		repo.NewCustomerRepo(pg),
	)
	petUseCase := usecase_pet.New(
		repo.NewPetRepo(pg),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, customerUseCase, petUseCase)
	return handler
}
