package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

var users []User
var nextID = 1

func main() {
	// Ginエンジンのインスタンスを作成
	router := gin.Default()

	// エンドポイントを定義
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/users", getUsers)
	router.GET("/user/:id", getUser)
	router.POST("/users", createUser)

	// サーバーを起動
	router.Run(":8080")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) //Atoi = ASCII to Integer

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	for _, user := range users { // _ → indexが不要の時
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "not existing id"})
}

func createUser(c *gin.Context) {
	var req CreateUserRequest

	// リクエストのJSONをreqに格納
	// エラーがあればerrに格納
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
		return
	}

	user := User{
		ID:   nextID,
		Name: req.Name,
	}

	nextID++
	users = append(users, user)

	c.JSON(http.StatusCreated, user)
}
