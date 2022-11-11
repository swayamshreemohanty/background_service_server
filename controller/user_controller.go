package controller

import (
	. "background_location_server/db"
	. "background_location_server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userMongoService UserMongoService
}

func UserNewController(userMongoService UserMongoService)  UserController{
	return UserController{userMongoService: userMongoService}
}



func (UserController *UserController)createUser(c *gin.Context)  {
	var newUser UserModel
	err:=c.ShouldBind(&newUser)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": false,
			"message":err.Error(),
			}) 
		return
	}else{
		newUser,err:=UserController.userMongoService.InsertUserDataToDB(&newUser)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status": false,
				"message":err.Error(),
				}) 
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"status": true,
			"message":"User added successfully",
			"user": newUser,
			}) 
		return
	}
}

func (UserController *UserController)updateUser(c *gin.Context)  {
	var existingUser UserLocationModel
	err:=c.ShouldBind(&existingUser)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": false,
			"message":err.Error(),
			}) 
		return
	}else{
		updatedData,username,err:=UserController.userMongoService.UpdateUserDataOnDB(&existingUser)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"status": false,
				"message":err.Error(),
				}) 
			return
		}else{
			c.JSON(http.StatusOK,gin.H{
				"status": true,
				"message":"User data updated successfully",
				"user": username,
				"data":updatedData,
				}) 
		}
		return
	}
}

func (UserController *UserController)RegisterUserRoutes(ginRouter *gin.RouterGroup)  {

	productRoute:=ginRouter.Group("users")
	productRoute.POST("create",UserController.createUser)
	productRoute.PUT("update",UserController.updateUser)

}