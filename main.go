package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-app/session"
	"go-app/users"
	"net/http"
)

type book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

var books = []book{
	{
		ID:     "1",
		Name:   "Data struct",
		Author: "Somchai",
		Price:  299.00,
	},
	{
		ID:     "2",
		Name:   "Algorithm",
		Author: "Somchai",
		Price:  299.00,
	},
	{
		ID:     "3",
		Name:   "Operating system",
		Author: "Silberschatz",
		Price:  399.00,
	},
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	paramID := c.Param("id")
	for _, book := range books {
		if book.ID == paramID {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
}

func addBook(c *gin.Context) {
	var newBook book

	err := c.BindJSON(&newBook)
	if err != nil {
		fmt.Println(err)
		return
	}

	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()

	redisStore, redisError := session.Init()
	if redisStore == nil {
		fmt.Println("Redis connect error", redisError)
		return
	}
	router.Use(sessions.Sessions("redis_session", redisStore))

	router.GET("/books", getBooks)
	router.GET("/book/:id", getBook)
	router.POST("/book", addBook)

	userRoute := router.Group("/user")
	userRoute.Use(users.UserMiddleware())
	users.RegisterRouter(userRoute)

	err := router.Run(":3000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
