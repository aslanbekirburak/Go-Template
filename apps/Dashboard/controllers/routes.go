package controllers

import (
	. "common_dashboard_backend/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type respMessage struct {
	Message    string `json:"message"`
	Email      string `json:"email"`
	ClientType string `json:"client_type"`
}

func InitRoutes(public, private *gin.RouterGroup) {

	dashboard := new(DashboardController)
	{
		public.POST("dashboardComponent/:namespace/:redisKey", dashboard.createDashboardComponent)
		public.GET("dashboardComponent/:namespace/:redisKey", dashboard.getDashboardComponent)
	}
}

func getUserIdFromToken(c *gin.Context) int {
	v, _ := c.Get("token-claims")
	claims := v.(jwt.MapClaims)
	userId := claims["userId"].(float64)
	return int(userId)
}

func generateSuccessResponse(data interface{}) map[string]interface{} {
	return gin.H{"data": data, "success": true, "errorCode": 0, "errorMessage": ""}
}

func generateFailResponse(err *ErrorType) map[string]interface{} {
	return gin.H{"data": nil, "success": false, "errorCode": err.Code, "errorMessage": err.Message}
}
