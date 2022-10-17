package http

import (
	"gitlab.com/golang-hse-2022/team1/tasks/internal/domain/models"
	"gitlab.com/golang-hse-2022/team1/tasks/pkg/infra/metrics"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/zaputil/zapctx"
)

func(a *Adapter) BindError(ctx *gin.Context, err error) {

	l := zapctx.Logger(ctx)
	l.Sugar().Errorf("request failed: %s", err.Error())

	span  := metrics.FromCtx(ctx)
	metrics.SetError(span, err)

	switch {
	case errors.Is(err, models.ErrForbidden):
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "access denied",
		})
	case errors.Is(err, models.ErrNotFound):
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "access denied",
		})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
