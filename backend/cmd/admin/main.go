package main

import (
	"backend/internal/core"
	"backend/internal/domain"
	"backend/pkg/db"
	"backend/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func main() {
	core.Inject().Invoke(func(app *ginx.App, dbe *db.DataBase, engine *gin.Engine,
		user *domain.User,
		role *domain.Role,
		team *domain.Team,
		address *domain.Address,
		product *domain.Product,
	) {
		go app.PProf()
		db.SetUp(dbe.Storage, user, role, team, address, product)
		app.Run(engine)
	})
}
