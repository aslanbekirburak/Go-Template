package controllers

import (
	// . "common_dashboard_backend/entities"

	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

func (dc DashboardController) createDashboardComponent(c *gin.Context) {
	redisKeyStr := c.Param("redisKey")
	body, err1 := ioutil.ReadAll(c.Request.Body)
	if err1 != nil {
		log.Fatal(err1)
		return
	}

	dashboardComponent := string(body)
	err := DashboardUseCase.CreateComponent(dashboardComponent, redisKeyStr)
	if err != nil {
		c.JSON(200, generateFailResponse(err))
		return
	}

	c.JSON(200, generateSuccessResponse("successfully created"))
	return
}

func (dc DashboardController) getDashboardComponent(c *gin.Context) {
	redisKeyStr := c.Param("redisKey")
	val, err := DashboardUseCase.GetComponent(redisKeyStr)
	if err != nil {
		c.JSON(200, generateFailResponse(err))
		return
	}
	c.JSON(200, generateSuccessResponse(val))
	return
}
