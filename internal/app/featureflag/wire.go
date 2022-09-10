//go:build wireinject
// +build wireinject

package featureflag

import (
	"github.com/google/wire"
	"snapp-featureflag/internal/app/featureflag/feature"
	"snapp-featureflag/internal/package/config"
	"snapp-featureflag/internal/package/service/redis"
)

func CreateApp() (*App, error) {
	panic(
		wire.Build(
			config.NewAppConfig,
			redis.NewCacheService,
			wire.Bind(new(redis.CacheService), new(redis.CacheServiceImpl)),
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
