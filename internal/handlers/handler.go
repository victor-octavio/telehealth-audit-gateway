package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-octavio/telehealth-audit-api/internal/handlers/diagnosis"
	"github.com/victor-octavio/telehealth-audit-api/internal/usecases"
)

type Handler struct {
	engine *gin.Engine
}

type Options struct {
	UC *usecases.Usecases
}

func New(opts Options) *Handler {
	engine := gin.Default()
	diagnosis.Init(opts.UC, engine.Group("/diagnosis"))

	return &Handler{
		engine: engine,
	}
}

func (h *Handler) ListenAndServe() error {
	return h.engine.Run(":8090")
}
