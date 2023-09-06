package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewFoodController),
	fx.Provide(NewMenuController),
	fx.Provide(NewOrderController),
)
