package bootstrap

import (
	"context"

	"go.uber.org/fx"
)

func WebApp() {
	//This line initializes a new application instance with the provided modules and options.
	app := fx.New(
		fx.Options(
			CommonModules,
		),
		fx.Invoke(startWebServer),
	)

	app.Run()
}

func startWebServer(lifecycle fx.Lifecycle) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.RunServer()
				return nil
			},
			OnStop: func(context context.Context) error {
				return nil
			},
		},
	)
}
