package bootstrap

import (
	"github.com/neerubhandari/restaurant-management/infrastructure"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	infrastructure.Module,
)
