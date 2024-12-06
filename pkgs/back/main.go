package main

import (
	"context"
	"database/sql"

	"server/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type PostAnnounceRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"  binding:"required"`
}

func main() {
	conn, err := sql.Open("postgres", "host=gpfdb port=5432 user=postgres password=password dbname=db sslmode=disable")
	if err != nil {
		panic(err)
	}
	queries := db.New(conn)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string {
			"http://localhost:8080",
		},
		AllowMethods: []string {
			"GET",
			"POST",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
	}))

	r.GET("/announces/get/", func(c *gin.Context) {
		if announces, err := queries.GetAnnounces(context.Background()); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"announces": announces,
			})
		}
	})
	r.POST("/announces/post/", func(c *gin.Context) {
		var req PostAnnounceRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		if res, err := queries.PostAnnounce(context.Background(), db.PostAnnounceParams{ Title: req.Title, Body: req.Body }); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"response": res,
			})
		}
	})

	r.Run(":63245")
}
