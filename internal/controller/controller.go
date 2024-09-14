package controller

import (
	"context"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mizmorr/rest-example/internal/model"
	"github.com/mizmorr/rest-example/service"
)

type UserController struct {
	ctx context.Context
	svc *service.UserWebService
}

func NewUsers(ctx context.Context, svc *service.UserWebService) *UserController {

	return &UserController{
		ctx: ctx,
		svc: svc,
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
//	@Success	200	{object}	model.User
//	@Failure	400	{object}	error
//	@Failure	404	{object}	error
//
// @Router		/user/{id} [get]
func (c *UserController) Get(g *gin.Context) {

	userid_raw, ok := g.Params.Get("id")
	if !ok {
		g.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"status": "id not provided"})
		return
	}
	userid, err := uuid.Parse(userid_raw)
	if err != nil {

		g.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"status": "could not parse userid"})
		return
	}
	user, err := c.svc.GetUser(g.Request.Context(), userid)

	if err != nil {

		g.AbortWithStatusJSON(http.StatusNotFound, map[string]string{"status": "could not found user"})
		return
	}
	g.JSON(http.StatusOK, user)

}

// Create		godoc
//
//	@Summary	Create user
//	@Tags		User
//	@Schemes
//	@Accept		json
//	@Produce	json
//	@Param		data body		model.UserCreateRequest	true	"user data"
//	@Success	200	{object}	model.User
//	@Failure	400	{object}	error
//	@Failure	304	{object}	error
//
// @Router		/user [post]
func (c *UserController) Create(g *gin.Context) {
	userCreateReq := model.UserCreateRequest{}
	err := g.Bind(&userCreateReq)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "create failed"})
		return
	}
	createdUser, err := c.svc.CreateUser(g.Request.Context(), &userCreateReq)

	if err != nil {
		g.AbortWithStatusJSON(http.StatusNotModified, gin.H{"status": "create failed"})
		return
	}

	g.JSON(http.StatusOK, createdUser)

}

// Delete			 godoc
//
//	@Summary	Delete user
//	@Tags		User
//	@Schemes
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"userid"
//	@Success	200	{object}	string
//	@Failure	400	{object}	error
//	@Failure	409	{object}	error
//
// @Router		/user/{id} [delete]
func (c *UserController) Delete(g *gin.Context) {

	userid_raw, ok := g.Params.Get("id")
	if !ok {
		g.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"status": "id not provided"})
		return
	}

	userid, err := uuid.Parse(userid_raw)
	if err != nil {

		g.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"status": "could not parse userid"})
		return
	}

	if err := c.svc.DeleteUser(g.Request.Context(), userid); err != nil {
		g.AbortWithStatusJSON(http.StatusConflict, gin.H{"status": "user not deleted"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("user with id %v was deleted", userid)})
}

// Update			 godoc
//
//	@Summary	Update user
//	@Tags		User
//	@Schemes
//	@Accept		json
//	@Produce	json
//	@Param		data body		model.UserUpdateRequest	true	"update user data"
//	@Success	200	{object}	string
//	@Failure	400	{object}	error
//	@Failure	409	{object}	error
//
// @Router		/user [put]
func (c *UserController) Update(g *gin.Context) {

	userUpdate := model.UserUpdateRequest{}

	err := g.Bind(&userUpdate)

	if err != nil {
		// g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "bad update request data"})
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": err})
		return
	}

	updUser, err := c.svc.UpdateUser(g.Request.Context(), &userUpdate)

	if err != nil {
		g.AbortWithStatusJSON(http.StatusNotModified, gin.H{"status": err})
		return
	}

	g.JSON(http.StatusOK, updUser)
}
