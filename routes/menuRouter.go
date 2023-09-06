package routes

import (
	"github.com/neerubhandari/restaurant-management/controllers"
	"github.com/neerubhandari/restaurant-management/infrastructure"
)

type MenuForRoutes struct {
	handler        infrastructure.Router
	menuController *controllers.MenuController
}

func NewMenuForRoutes(
	handler infrastructure.Router,
	menuController *controllers.MenuController,
) *MenuForRoutes {
	return &MenuForRoutes{
		handler:        handler,
		menuController: menuController,
	}
}

func (mr *MenuForRoutes) Setup() {
	api := mr.handler.Group("/menus")
	api.GET("", mr.menuController.GetMenus)
	api.GET("/:menu_id", mr.menuController.GetMenu)
	api.POST("", mr.menuController.CreateMenu)
	api.PATCH("/:menu_id", mr.menuController.UpdateMenu)
}
