package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/joaomarcosg/Projeto-Taskify/internal/services"
)

type Application struct {
	Router      *chi.Mux
	TaskService services.TaskService
}
