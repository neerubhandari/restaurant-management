package bootstrap

import (
	"github.com/neerubhandari/restaurant-management/controllers"
	"github.com/neerubhandari/restaurant-management/infrastructure"
	"github.com/neerubhandari/restaurant-management/routes"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	infrastructure.Module,
	routes.Module,
	controllers.Module,
	controllers.MigrateModule,
)
