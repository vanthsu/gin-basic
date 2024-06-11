package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vanthsu/gin-basic/models"
)

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			formatResponse("GetAllUsers() error", err.Error(), nil),
		)
		return
	}

	context.JSON(
		http.StatusOK,
		formatResponse("ok", "", users),
	)
}

func getUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			formatResponse("Invalid :id input", err.Error(), nil),
		)
		return
	}

	user, err := models.GetUserById(id)
	if err != nil {
		context.JSON(
			http.StatusNotFound,
			formatResponse("Cannot find user", err.Error(), nil),
		)
		return
	}

	context.JSON(
		http.StatusOK,
		formatResponse("ok", "", user),
	)
}

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			formatResponse("context.ShouldBindJSON() error", err.Error(), nil),
		)
		return
	}

	_, err = user.Save()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			formatResponse("user.Save() error", err.Error(), nil),
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		formatResponse("user created successfully", "", user),
	)
}

func updateUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			formatResponse("Invalid :id input", err.Error(), nil),
		)
		return
	}

	_, err = models.GetUserById(id)
	if err != nil {
		context.JSON(
			http.StatusNotFound,
			formatResponse("Cannot find user", err.Error(), nil),
		)
		return
	}

	var updatedUser models.User
	err = context.ShouldBindJSON(&updatedUser)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			formatResponse("context.ShouldBindJSON() error", err.Error(), nil),
		)
		return
	}

	updatedUser.ID = id
	_, err = updatedUser.Update()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			formatResponse("updatedUser.Update() error", err.Error(), nil),
		)
	}

	context.JSON(
		http.StatusOK,
		formatResponse("User updated successfully", "", updatedUser),
	)
}

func deleteUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			formatResponse("Invalid :id input", err.Error(), nil),
		)
		return
	}

	user, err := models.GetUserById(id)
	if err != nil {
		context.JSON(
			http.StatusNotFound,
			formatResponse("Cannot find user", err.Error(), nil),
		)
		return
	}

	err = user.Delete()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			formatResponse("User delete failed", err.Error(), nil),
		)
		return
	}

	context.JSON(
		http.StatusOK,
		formatResponse("User deleted succussfully", "", nil),
	)
}

func formatResponse(msg string, err string, data any) *gin.H {
	return &gin.H{
		"message": msg,
		"error":   err,
		"data":    data,
	}
}
