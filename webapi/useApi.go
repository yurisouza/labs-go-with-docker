package webapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yurisouza/labs-go-with-docker/models"
	"github.com/yurisouza/labs-go-with-docker/services"
)

func serveUsuarioResource(engine *gin.Engine) {
	v1 := engine.Group("/api/v1/users")
	{
		v1.GET("/", getAllUsers)
		v1.GET("/:id", getUser)
		v1.POST("/", insertUser)
		v1.PUT("/:id", updateUser)
		v1.DELETE("/:id", removeUser)
	}
	base := engine.Group("/api/create")
	{
		base.GET("/", createStructure)
	}
}

func getAllUsers(c *gin.Context) {

	users, err := services.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro: " + err.Error()})
	} else if users != nil {
		c.JSON(http.StatusOK, &users)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
	}

}

func getUser(c *gin.Context) {
	id := c.Param("id")

	user, err := services.GetUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro: " + err.Error()})
	} else if user != nil {
		c.JSON(http.StatusOK, &user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
	}
}

func insertUser(c *gin.Context) {
	user := models.User{}

	body, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	errDb := services.InsertUser(&user)

	if errDb != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro: " + err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "User created with success"})
	}
}

func updateUser(c *gin.Context) {
	id := c.Param("id")

	user := models.User{}

	body, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	userBd, errDb := services.GetUser(id)

	if userBd != nil && errDb == nil {
		errDb := services.UpdateUser(&user)

		if errDb != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro: " + errDb.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "User update with success", "user": &user})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
	}

}

func removeUser(c *gin.Context) {
	id := c.Param("id")

	user, err := services.GetUser(id)

	if user != nil && err == nil {
		err := services.RemoveUser(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro: " + err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "User removed with success"})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
	}

}

func createStructure(c *gin.Context) {
	err := services.CreateStructure()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro: " + err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Table created with success"})
	}
}
