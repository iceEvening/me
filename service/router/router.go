package router

import (
	"github.com/iceEvening/me/service/middleware/jwt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/iceEvening/me/service/model"
)

//Router is a struct which manages routers in different cases.
type Router struct {
	E     *gin.Engine
	model *model.Model
}

//GetRouter returns a Router struct pointer
func GetRouter(e *gin.Engine, db *gorm.DB) *Router {
	return &Router{
		E:     e,
		model: model.GetModel(db),
	}
}

func Atoi(c *gin.Context, a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "invalid id",
		})
	}
	return i
}

func CheckEditAuth(c *gin.Context, ID uint) bool {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "internal error",
		})
	}
	if claims.(*jwt.CustomClaims).ID != ID {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "invalid id",
		})
	}
	return true
}
