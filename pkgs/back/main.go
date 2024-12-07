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

type PostUserRequest struct {
	Name string `json:"name" binding:"required"`
}

type BanOrUnbanUserRequest struct {
	ID     int32 `json:"id"     binding:"required"`
	Banned *bool `json:"banned" binding:"required"`
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
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
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

	r.GET("/users/get/", func(c *gin.Context) {
		if users, err := queries.GetUsers(context.Background()); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"users": users,
			})
		}
	})
	r.POST("/users/post/", func(c *gin.Context) {
		var req PostUserRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		if res, err := queries.PostUser(context.Background(), req.Name); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"response": res,
			})
		}
	})
	r.POST("/users/ban-or-unban/", func(c *gin.Context) {
		var req BanOrUnbanUserRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		if res, err := queries.BanOrUnbanUser(context.Background(), db.BanOrUnbanUserParams{ ID: req.ID, Banned: *req.Banned }); err != nil {
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
