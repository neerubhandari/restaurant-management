package main

import (
	"github.com/neerubhandari/restaurant-management/bootstrap"
	"github.com/neerubhandari/restaurant-management/utils"
)

func main() {

	utils.LoadEnv()
	bootstrap.WebApp()

}
