package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CacheController struct {
	cache interface {
		Take(ctx context.Context) interface{}
	}
}

func NewCache(i interface {
	Take(ctx context.Context) interface{}
}) *CacheController {
	return &CacheController{
		cache: i,
	}
}

// Get			 godoc
//
//	@Summary	Get cache
//	@Tags		cache
//	@Schemes
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	interface{}
//
// @Router		/cache [get]
func (cc *CacheController) Get(g *gin.Context) {

	data := cc.cache.Take(g.Request.Context())

	g.JSON(http.StatusOK, data)
}
