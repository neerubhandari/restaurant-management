package controllers

import "go.uber.org/fx"

type Migrater interface {
	Migrate() error
}

type controllers []Migrater

var MigrateModule = fx.Options(
	fx.Provide(NewMigrate),
	fx.Invoke(Migrate),
)

func NewMigrate(orderContoller *OrderController,

// menuController *MenuController,
// foodController *FoodController,
) controllers {
	return controllers{orderContoller}// menuController,
	//  foodController

}
func Migrate(c controllers) error {
	for _, repo := range c {
		if err := repo.Migrate(); err != nil {
			return err
		}
	}
	return nil
}
