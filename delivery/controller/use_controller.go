package controller

import (
	"SplitAll/delivery/api"
	"SplitAll/model"
	"SplitAll/usecase"
	"SplitAll/utils"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router    *gin.RouterGroup
	routerDev *gin.RouterGroup
	ucUser    usecase.UserUsecase
	api.BaseApi
}

func (u *UserController) GetOcrData(c *gin.Context) {
	// Retrieve the file
	file, err := c.FormFile("file")
	if err != nil {
		u.Failed(c, utils.UploadImageError())
		return
	}

	// Validate file size (between 10KB and 2MB)
	if file.Size < 10*1024 || file.Size > 2*1024*1024 {
		u.Failed(c, utils.UploadImageFileLimitation())
		return
	}

	// Validate file format (only accept jpg/jpeg)
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" {
		u.Failed(c, utils.UploadImageTypeError())
		return
	}

	// Call the usecase to handle the OCR process
	imageInfo, err := u.ucUser.GetOcrInfo(file)
	if err != nil {
		fmt.Println("Error in use_controller --> ", err) // Useful for debugging
		u.Failed(c, utils.UploadImageError())
		return
	}

	// Success response
	u.Success(c, imageInfo, 200, "")
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

	router.POST("/image_ocr", controller.GetOcrData)

	// router.POST("/image", controller.UploadImage)

	router.POST("/convertDataToText", controller.RecepientSend)

	return &controller
}
