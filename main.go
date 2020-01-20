package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

type errResTmpl struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	r.LoadHTMLGlob("resources/templates/*.html")
	r.Static("/assets", "./resources/assets")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/api/new", func(c *gin.Context) {
		costStr := c.DefaultQuery("cost", "10")
		cost, err := strconv.Atoi(costStr)
		if err != nil {
			c.JSON(400, errResTmpl{400, "Bad request", "please specify cost query by an integer"})
			return
		}
		h, err := bcrypt.GenerateFromPassword([]byte("河野太郎"), cost)
		if err != nil {
			c.JSON(500, errResTmpl{500, "Internal server error", err.Error()})
			return
		}
		c.JSON(200, gin.H{"hashed": string(h)})
	})
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
