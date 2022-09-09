//go:build wireinject
// +build wireinject

package featureflag

import (
	"github.com/google/wire"
	"snapp-featureflag/internal/package/config"
)

func CreateApp() (*App, error) {
	panic(
		wire.Build(
			config.NewAppConfig,
			NewHttpServeMux,
			NewApp,
		),
	)
}
