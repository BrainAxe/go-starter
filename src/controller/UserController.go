package controller

import (
	"go-starter/src/dao"
	db "go-starter/src/db/sqlc"
	"go-starter/src/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
}

// Create new user
func CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, hashErr := helper.HashPassword(req.Password)

	if hashErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": hashErr.Error(),
		})
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Mobile:   req.Mobile,
	}

	_, err := dao.SqlSession.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  "success",
		"data": "registration successful",
	})

}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// Get user list
func GetUserList(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	userList, err := dao.SqlSession.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": userList,
	})
}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type updateUserRequest struct {
	Email  string `json:"email" binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
}

// Update user
func UpdateUser(ctx *gin.Context) {
	var reqOne getUserRequest
	var reqTwo updateUserRequest
	if err := ctx.ShouldBindUri(&reqOne); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&reqTwo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	arg := db.UpdateUserParams{
		ID:     reqOne.ID,
		Email:  reqTwo.Email,
		Mobile: reqTwo.Mobile,
	}

	if user, err := dao.SqlSession.UpdateUser(ctx, arg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "successfully updated",
			"data": user,
		})
	}
}

// Delete user
func DeleteUserById(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dao.SqlSession.DeleteUser(ctx, req.ID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "successfully deleted",
			"data": req.ID,
		})
	}
}
