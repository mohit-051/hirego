package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohit-051/hirego/internal/core/domain"
	"github.com/mohit-051/hirego/internal/core/ports"
	"github.com/mohit-051/hirego/internal/helper"
)

var HRHandler *hrHandler

type hrHandler struct {
	hrService ports.HRService
}

func (h *hrHandler) Initialize(hrserv ports.HRService) {
	HRHandler = &hrHandler{
		hrService: hrserv,
	}
}

func (h *hrHandler) Signup(c *gin.Context) {
	var hr domain.HrManager
	c.BindJSON(&hr)

	// Call the signup service to signup the hr.
	message, err := h.hrService.Signup(hr)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)
}

func (h *hrHandler) Login(c *gin.Context) {
	// The BindJSON is used to extract the JSON data from the request body.
	var hr domain.HrManager
	c.BindJSON(&hr)

	// Call the Login service to signup the hr.
	message, err := h.hrService.Login(hr)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)
}

func (h *hrHandler) SetHrProfileInformation(c *gin.Context) {

	var hrInformation domain.HrProfileInformation
	c.BindJSON(&hrInformation)

	// Update the hr's email based on the hr_email from the context
	hrMap := c.MustGet("username").(map[string]interface{})

	// Call the service to update the hr's profile information.
	message, err := h.hrService.SetHrProfileInformation(hrMap["username"].(string), hrInformation)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)
}

func (h *hrHandler) GetProfileInformation(c *gin.Context) {
	// Get the hr's email from the context
	hrMap := c.MustGet("username").(map[string]interface{})

	// Call the service to get the hr's profile information.
	hrInformation, err := h.hrService.GetProfileInformation(hrMap["username"].(string))

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, hrInformation, nil)
}

func (h *hrHandler) JobPosting(c *gin.Context) {
	var jobPosting domain.JobPosting
	c.BindJSON(&jobPosting)

	jobPosting.Username = c.MustGet("username").(map[string]interface{})["username"].(string)

	// Call the service to post the job.
	message, err := h.hrService.JobPosting(jobPosting)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)
}

func (h *hrHandler) GetJobPosting(c *gin.Context) {
	// Get the hr's email from the context
	// hrMap := c.MustGet("username").(map[string]interface{})
	var jobPosting domain.JobPosting
	c.BindJSON(&jobPosting)

	// Call the service to get the hr's profile information.
	jobPosting, err := h.hrService.GetJobPosting(jobPosting.JobID)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, jobPosting, nil)
}

func (h *hrHandler) GetAllJobPosting(c *gin.Context) {

	// Call the service to get the hr's profile information.
	jobPosting, err := h.hrService.GetAllJobPosting()

	fmt.Println(jobPosting)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, jobPosting, nil)
}

func (h *hrHandler) HrSpecificJobPosting(c *gin.Context) {
	// Get the hr's email from the context
	hrMap := c.MustGet("username").(map[string]interface{})

	// Call the service to get the hr's profile information.
	jobPosting, err := h.hrService.HrSpecificJobPosting(hrMap["username"].(string))

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, jobPosting, nil)
}

func (h *hrHandler) GetAllApplicants(c *gin.Context) {
	// Get the hr's email from the context
	hrMap := c.MustGet("username").(map[string]interface{})

	fmt.Println(hrMap)
	var jobId domain.JobPosting
	c.BindJSON(&jobId)

	// Call the service to get the hr's profile information.
	applicants, err := h.hrService.GetAllApplicants(jobId.JobID)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, applicants, nil)
}
