package app

import (
	"github.com/jupripratama/go-homework/config"
)

type Application struct {
	Config *config.Config
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
	}

	return application
}
