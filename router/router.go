package router

import (
	"strconv"

	"github.com/artfunder/users-service/repository"
	"github.com/artfunder/users-service/service"
	"github.com/gin-gonic/gin"
)

// CreateRouter ...
func CreateRouter(s *service.UsersService) *gin.Engine {
	r := gin.Default()

	r.GET("/:id", getOneHandler(s))

	return r
}

func getOneHandler(s *service.UsersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"error":   "id must be number",
			})
			return
		}

		user, err := s.GetOne(id)
		if err == repository.ErrNoUserWithID {
			c.JSON(404, gin.H{
				"success": false,
				"error":   "No user with id " + idStr,
			})
			return
		}

		c.JSON(200, gin.H{
			"success": true,
			"user":    user,
		})
	}
}
