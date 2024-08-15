package controller

import (
	"SplitAll/delivery/api"
	"SplitAll/model"
	"SplitAll/usecase"
	"SplitAll/utils"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router    *gin.RouterGroup
	routerDev *gin.RouterGroup
	ucUser    usecase.UserUsecase
	api.BaseApi
}

func (u *UserController) RecepientSend(c *gin.Context) {

	var bodyRequest []model.UserRecepient

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	data, err := u.ucUser.UserSendRecepeint(bodyRequest)

	if err != nil {
		u.Failed(c, err)
		return
	}

	detailMsg := "Picture Successfully Receive"
	u.Success(c, data, detailMsg, "")

}

func (u *UserController) UploadImage(c *gin.Context) {
	// Retrieve the file
	file, err := c.FormFile("file")
	if err != nil {
		u.Failed(c, utils.UploadImageError())
		return
	}

	// Check file size
	if file.Size < 10*1024 || file.Size > 2*1024*1024 {
		u.Failed(c, utils.UploadImageFileLimitation())
		return
	}

	// Check file format
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" {
		u.Failed(c, utils.UploadImageTypeError())
		return
	}

	imageUrl, err := u.ucUser.SaveImageURL(file)
	if err != nil {
		u.Failed(c, utils.UploadImageError())
		return
	}

	detailMsg := 200
	u.Success(c, imageUrl, detailMsg, "")
}

func NewUserController(router *gin.RouterGroup, routerDev *gin.RouterGroup, ucUser usecase.UserUsecase) *UserController {
	controller := UserController{
		router:    router,
		routerDev: routerDev,
		ucUser:    ucUser,
		BaseApi:   api.BaseApi{},
	}

	router.POST("/image", controller.UploadImage)

	router.POST("/convertDataToText", controller.RecepientSend)

	return &controller
}
