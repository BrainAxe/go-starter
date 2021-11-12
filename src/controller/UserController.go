package controller

import (
	"go-starter/src/model"
	"go-starter/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create new user
func CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := service.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg":  "success",
		})
	}
}

// Get user list
func GetUserList(c *gin.Context) {
	userList, err := service.GetAlluser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": userList,
		})
	}
}

// Update user
func UpdateUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	user, err := service.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&user)
	// fmt.Println(user)
	if err = service.UpdateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "successfully updated",
			"data": user,
		})
	}
}

// Delete user
func DeleteUserById(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	} else if err := service.DeleteUserById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "successfully deleted",
			"data": id,
		})
	}
}
