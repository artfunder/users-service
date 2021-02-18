package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/artfunder/structs"
	"github.com/artfunder/users-service/repository"
	"github.com/artfunder/users-service/service"
	"github.com/gin-gonic/gin"
)

// CreateRouter ...
func CreateRouter(s *service.UsersService) http.Handler {
	// r := gin.Default()

	// r.GET("/:id", getOneHandler(s))
	// r.POST("/", createHandler(s))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CLIENT_SITE"))
	})
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

func createHandler(s *service.UsersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInfo structs.User
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"error":   "Bad Request",
			})
			return
		}
		err = json.Unmarshal(body, &userInfo)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"error":   "Bad JSON",
			})
			return
		}

		user, err := s.Create(userInfo)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		c.JSON(201, gin.H{
			"success": true,
			"user":    user,
		})
	}
}
