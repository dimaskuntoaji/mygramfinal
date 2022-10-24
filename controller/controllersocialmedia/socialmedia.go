package controllersocialmedia

import (
	"net/http"
	"strconv"

	"finall/helper"
	"finall/model/modelsocialmedia"
	"finall/service/servicesocialmedia"

	"github.com/gin-gonic/gin"
)

type ControllerSocialMedia interface {
	Create(ctx *gin.Context)
	GetList(ctx *gin.Context)
	UpdateByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type controller struct {
	srv servicesocialmedia.ServiceSocialMedia
}

func New(srv servicesocialmedia.ServiceSocialMedia) ControllerSocialMedia {
	return &controller{srv}
}


func (c *controller) Create(ctx *gin.Context) {
	data := new(modelsocialmedia.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	data.UserID = ctx.MustGet("user_id").(uint)

	response, err := c.srv.Create(*data)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, response, nil))
}

func (c *controller) GetList(ctx *gin.Context) {

	response, err := c.srv.GetList()
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

func (c *controller) UpdateByID(ctx *gin.Context) {
	data := new(modelsocialmedia.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	idString := ctx.Param("socialmediaid")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	data.ID = uint(id)

	response, err := c.srv.UpdateByID(*data)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

func (c *controller) DeleteByID(ctx *gin.Context) {
	idString := ctx.Param("socialmediaid")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	err = c.srv.DeleteByID(uint(id))
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, map[string]interface{}{"message": "Your social media has been successfully deleted"}, nil))
}


