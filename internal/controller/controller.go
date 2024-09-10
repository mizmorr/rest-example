package controller

import (
	"context"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mizmorr/rest-example/pkg/logger"
	"github.com/mizmorr/rest-example/service"
)

type UserController struct {
	ctx    context.Context
	svc    *service.UserWebService
	logger *logger.Logger
}

func NewUsers(ctx context.Context, svc *service.UserWebService, logger *logger.Logger) *UserController {

	return &UserController{
		ctx:    ctx,
		svc:    svc,
		logger: logger,
	}
}

// Get			 godoc
//
//	@Summary	Get user
//	@Tags		User
//	@Schemes
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"userid"
//	@Success	200	{object}	user.User
//	@Failure	400	{object}	error
//	@Router		/user/{id} [get]
func (c *UserController) Get(g *gin.Context) {

	userid_raw, ok := g.Params.Get("id")
	if !ok {
		g.AbortWithError(http.StatusBadRequest, fmt.Errorf("no id provided"))
	}
	userid, err := uuid.Parse(userid_raw)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, fmt.Errorf("could not parse userid: %v", userid_raw))
	}
	user, err := c.svc.GetUser(g.Request.Context(), userid)

	if err != nil {
		g.AbortWithError(http.StatusNotFound, fmt.Errorf("could not found user: %v", userid))
	}
	g.JSON(http.StatusOK, user)

}
