package wire

import (
	"go-29/internal/adaptor"
	"go-29/internal/data/repository"
	"go-29/internal/usecase"
	"go-29/pkg/middleware"
	"go-29/pkg/utils"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Wiring(repo repository.Repository, mLogger middleware.LoggerMiddleware, logger *zap.Logger, config utils.Configuration) *chi.Mux {
	router := chi.NewRouter()
	router.Use(mLogger.LoggingMiddleware)
	rV1 := chi.NewRouter()
	wireUser(rV1, repo, logger, config)
	router.Mount("/api/v1", rV1)

	return router
}

func wireUser(router *chi.Mux, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecaseUser := usecase.NewUserService(repo, logger, config)
	adaptorUser := adaptor.NewHandlerUser(usecaseUser, logger)
	router.Post("/register", adaptorUser.Register)
}
