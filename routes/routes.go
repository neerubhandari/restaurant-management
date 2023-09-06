package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewFoodForRoutes),
	fx.Provide(NewMenuForRoutes),
	fx.Invoke(Setup),
)

// Routes contains multiple routes
type Routers []Router
type Router interface {
	Setup()
}

func NewRoutes(foodRoutes *FoodForRoutes, menuRoutes *MenuForRoutes) Routers {
	return Routers{foodRoutes, menuRoutes}
}

// Setup all the route
func Setup(r Routers) {

	for _, route := range r {
		route.Setup()
	}

}
