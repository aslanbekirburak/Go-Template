package controllers

import (
	// . "common_dashboard_backend/entities"

	"common_dashboard_backend/entities"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

func (dc DashboardController) createDashboardComponent(c *gin.Context) {
	redisKeyStr := c.Param("redisKey")
	namespace := c.Param("namespace")

	body, err1 := ioutil.ReadAll(c.Request.Body)
	if err1 != nil {
		log.Fatal(err1)
		return
	}

	dashboardComponent := string(body)
	err := DashboardUseCase.CreateComponent(dashboardComponent, namespace, namespace+"_"+redisKeyStr)
	if err != nil {
		c.JSON(200, generateFailResponse(err))
		return
	}

	c.JSON(200, generateSuccessResponse("successfully created"))
	return
}

func (dc DashboardController) getDashboardComponent(c *gin.Context) {
	redisKeyStr := c.Param("redisKey")
	namespace := c.Param("namespace")

	val, err := DashboardUseCase.GetComponent(namespace + "_" + redisKeyStr)
	if err != nil {
		c.JSON(200, generateFailResponse(err))
		return
	}
	if val == "" {
		c.JSON(404, generateFailResponse(&entities.ErrorType{
			Code:    404,
			Message: "key not found",
		}))
		return
	}
	c.JSON(200, generateSuccessResponse(val))
	return
}

func (dc DashboardController) getDashboardNamespaces(c *gin.Context) {

	val, err := DashboardUseCase.GetNamespace()
	if err != nil {
		c.JSON(200, generateFailResponse(err))
		return
	}
	c.JSON(200, generateSuccessResponse(val))
	return
}

func (dc DashboardController) getDashboardKeys(c *gin.Context) {
	redisKey := c.Param("redisKey")

	val, err := DashboardUseCase.GetNameKeys(redisKey)
	if err != nil {
		c.JSON(200, generateFailResponse(err))
		return
	}
	c.JSON(200, generateSuccessResponse(val))
	return
}

func (dc DashboardController) deleteDashboardComponent(c *gin.Context) {
	redisKey := c.Param("redisKey")
	namespace := c.Param("namespace")

	err := DashboardUseCase.DeleteComponent(namespace, redisKey)
	if err != nil {
		c.JSON(200, generateFailResponse(err))
		return
	}
	c.JSON(200, generateSuccessResponse(&entities.ErrorType{
		Code:    200,
		Message: "Successfully Deleted",
	}))
	return
}
