package main

import (
	"log"

	"real-estate-agency-onion/internal/application/usecases/agent"
	"real-estate-agency-onion/internal/infrastructure/auth"
	"real-estate-agency-onion/internal/infrastructure/config"
	"real-estate-agency-onion/internal/infrastructure/logger"
	"real-estate-agency-onion/internal/infrastructure/persistence/postgres"
	repos "real-estate-agency-onion/internal/infrastructure/persistence/postgres/repositories"
	"real-estate-agency-onion/internal/interfaces/http/handlers"
	"real-estate-agency-onion/internal/interfaces/http/routes"
)

func main() {
	cfg := config.Load()
	logg := logger.New()

	db, err := postgres.NewDB(cfg.DatabaseURL)
	if err != nil {
		logg.Error("failed to connect to database", "error", err)
		log.Fatal(err)
	}

	tm := auth.NewJWTTokenManager(cfg.JWTSecret, cfg.JWTExpiration)

	agentRepo := repos.NewAgentRepository(db)

	createUC := agent.NewCreateUseCase(agentRepo)
	getByIDUC := agent.NewGetAgentByIDUseCase(agentRepo)
	listUC := agent.NewListUseCase(agentRepo)
	updateUC := agent.NewUpdateUseCase(agentRepo)
	softDeleteUC := agent.NewSoftDeleteUseCase(agentRepo)
	deactivateUC := agent.NewDeactivateUseCase(agentRepo)

	agentH := handlers.NewAgentHandler(createUC, getByIDUC, listUC, updateUC, softDeleteUC, deactivateUC)

	r := routes.Setup(agentH, tm)

	logg.Info("starting server", "port", cfg.ServerPort)
	if err := r.Run(cfg.ServerPort); err != nil {
		logg.Error("server failed", "error", err)
		log.Fatal(err)
	}
}
