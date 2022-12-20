package http

import (
	"time"

	"gitlab.com/golang-hse-2022/team1/tasks/pkg/infra/logger"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/juju/zaputil/zapctx"
)

func initRouter(a *Adapter, r *gin.Engine, l logger.Logger) {
	r.Use(func(ctx *gin.Context) {
		lCtx := zapctx.WithLogger(ctx.Request.Context(), l)
		ctx.Request = ctx.Request.WithContext(lCtx)
	})
	r.Use(ginzap.Ginzap(l, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(l, true))

	authorized := r.Group("/tasks/api/v1/tasks/")
	authorized.Use(a.AuthRequired())
	{
		authorized.GET("/", a.getTasks)
		authorized.GET("/:taskId", a.getDescription)
		authorized.POST("/approve/:taskId", a.approve)
		authorized.POST("/decline/:taskId", a.decline)
		authorized.DELETE("/:taskId", a.delete)
		limited := authorized.Group("/")
		limited.Use(a.Limited())
		{
			limited.POST("/", a.create)
		}
	}
}
