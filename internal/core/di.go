package core

import (
	"backend/internal/app/admin"
	"backend/internal/domain"
	"backend/internal/logic"
	"backend/internal/pkg/config"
	"backend/internal/repository"
	"backend/internal/router"
	"backend/pkg/cache"
	"backend/pkg/db"
	"backend/pkg/logx"
	"backend/pkg/token"
	"go.uber.org/dig"
)

func Inject() *dig.Container {
	container := dig.New()

	container.Provide(config.NewConfig)
	container.Provide(config.New)

	container.Provide(db.New)
	container.Provide(logx.New)
	container.Provide(cache.New)

	container.Provide(NewApp)
	container.Provide(NewLogConfig)
	container.Provide(NewMySQLConfig)
	container.Provide(NewRedisConfig)
	container.Provide(NewJWTConfig)

	container.Provide(token.New)

	container.Provide(repository.NewUserRepository)
	container.Provide(logic.NewUserLogic)
	container.Provide(admin.NewUserController)

	container.Provide(domain.NewUser)
	container.Provide(domain.NewRole)
	container.Provide(domain.NewTeam)
	container.Provide(domain.NewAddress)
	container.Provide(domain.NewProduct)

	container.Provide(repository.NewUserRepositoryFace)
	container.Provide(logic.NewUserLogicFace)
	container.Provide(cache.NewRedisFace)
	container.Provide(token.NewJsonWebTokenFace)

	container.Provide(router.SetUpRouter)
	return container
}
