package main

import (
	"fmt"

	"github.com/lovelyrrg51/go_backend/app/config"
	"github.com/lovelyrrg51/go_backend/app/routes"
)

func main() {
	port := config.Cfg.AppPort
	runPort := fmt.Sprint(":", port)

	routes.Route.Run(runPort)
}
