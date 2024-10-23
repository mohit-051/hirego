package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohit-051/hirego/internal/core/domain"
	"github.com/mohit-051/hirego/internal/core/ports"
	"github.com/mohit-051/hirego/internal/helper"
)

var UserHandler *userHandler

type userHandler struct {
	// *base
	userService ports.UserService
}

func (h *userHandler) Initialize(userserv ports.UserService) {
	UserHandler = &userHandler{
		userService: userserv,
	}
}

func (h *userHandler) Signup(c *gin.Context) {
	var user domain.User
	c.BindJSON(&user)

	// Call the signup service to signup the user.
	message, err := h.userService.Signup(user)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)

}

func (h *userHandler) Login(c *gin.Context) {
	// The BindJSON is used to extract the JSON data from the request body.
	var user domain.User
	c.BindJSON(&user)

	// Call the Login service to signup the user.
	message, err := h.userService.Login(user)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)
}

func (h *userHandler) SetUserWrorkInformation(c *gin.Context) {

	var workinformation domain.WorkInformation
	c.BindJSON(&workinformation)

	// Update the user's email based on the user_email from the context
	userMap := c.MustGet("username").(map[string]interface{})

	// Call the service to save the user work information.
	message, err := h.userService.SetUserWrorkInformation(userMap["username"].(string), workinformation)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)
}

func (h *userHandler) GetUserWorkInformation(c *gin.Context) {

	var username domain.UserName
	c.BindJSON(&username)

	// Call the signup service to signup the user.
	message, err := h.userService.GetUserWorkInformation(username.Username)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)
}

func (h *userHandler) SetUserProfileInformation(c *gin.Context) {
	// The BindJSON is used to extract the JSON data from the request body.
	var userInformation domain.UserInformation
	c.BindJSON(&userInformation)

	// Update the user's email based on the user_email from the context
	userMap := c.MustGet("username").(map[string]interface{})

	// Call the Profile service to save the user profile.
	message, err := h.userService.SetUserProfileInformation(userMap["username"].(string), userInformation)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)
}

func (h *userHandler) GetProfileInformation(c *gin.Context) {
	// Update the user's email based on the user_email from the context
	userMap := c.MustGet("username").(map[string]interface{})

	// Call the service to get the user profile information.
	message, err := h.userService.GetProfileInformation(userMap["username"].(string))

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)
}

func (s *userHandler) ApplyForJobPosting(c *gin.Context) {

	var jobPosting domain.JobApplication
	c.BindJSON(&jobPosting)

	userMap := c.MustGet("username").(map[string]interface{})

	jobPosting.Username = userMap["username"].(string)

	// Call the login repo to insert the data of the user.
	message, err := s.userService.ApplyForJobPosting(jobPosting)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.Response(c, 200, message, nil)

}

func (s *userHandler) GetAppliedJobPosting(c *gin.Context) {

	userMap := c.MustGet("username").(map[string]interface{})

	// Call the login repo to insert the data of the user.
	message, err := s.userService.GetAppliedJobPosting(userMap["username"].(string))

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)

}
