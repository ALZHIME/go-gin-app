package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-app/common"
	"go-app/session"
	"go-app/statics"
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
	router.Use(common.EnableCORS())

	redisStore, redisError := session.Init()
	if redisStore == nil {
		fmt.Println("Redis connect error", redisError)
		return
	}
	router.Use(sessions.Sessions("cookie_name", redisStore))

	router.GET("/books", getBooks)
	router.GET("/book/:id", getBook)
	router.POST("/book", addBook)

	userRoute := router.Group("/user")
	userRoute.Use(users.UserMiddleware())
	users.RegisterRouter(userRoute)

	protectedStaticRoute := router.Group("/")
	protectedStaticRoute.Use(statics.ProtectedMiddleware())
	statics.RegisterProtected(protectedStaticRoute)

	publicStaticRoute := router.Group("/")
	statics.RegisterPublic(publicStaticRoute)

	// https://stackoverflow.com/questions/10175812/how-to-generate-a-self-signed-ssl-certificate-using-openssl
	err := router.RunTLS(":3000", "./cert/localhost.crt", "./cert/localhost.key")
	if err != nil {
		fmt.Println(err)
		return
	}
}
