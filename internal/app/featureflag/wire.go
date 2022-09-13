//go:build wireinject
// +build wireinject

package featureflag

import (
	"github.com/google/wire"
	"snapp-featureflag/internal/app/featureflag/feature"
	"snapp-featureflag/internal/package/config"
	"snapp-featureflag/internal/package/service/cache"
)

func CreateApp() (*App, error) {
	panic(
		wire.Build(
			config.NewAppConfig,
			cache.NewCacheService,
			wire.Bind(new(cache.Service), new(cache.ServiceImpl)),
			feature.NewRepository,
			wire.Bind(new(feature.Repository), new(feature.RepositoryImpl)),
			feature.NewCommandHandler,
			feature.NewQueryHandler,
			NewApiHandler,
			NewHttpServeMux,
			NewApp,
		),
	)
}
