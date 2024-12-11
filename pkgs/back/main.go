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

type PostInquiryRequest struct {
	Usrid int32  `json:"usrid" binding:"required"`
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"  binding:"required"`
}

type ReplyInquiryRequest struct {
	ID    int32  `json:"id"    binding:"required"`
	Reply string `json:"reply" binding:"required"`
}

type PostItemRequest struct {
	Name string `json:"name" binding:"required"`
}

type PostItemToUserRequest struct {
	Usrid  int32 `json:"usrid"  binding:"required"`
	Itmid  int32 `json:"itmid"  binding:"required"`
	Amount int32 `json:"amount" binding:"required"`
}

type PostItemToAllUsersRequest struct {
	Itmid  int32 `json:"itmid"  binding:"required"`
	Amount int32 `json:"amount" binding:"required"`
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

	r.GET("/inquiries/get/", func(c *gin.Context) {
		if inquiries, err := queries.GetInquiries(context.Background()); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"inquiries": inquiries,
			})
		}
	})
	r.POST("/inquiries/post/", func(c *gin.Context) {
		var req PostInquiryRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		if res, err := queries.PostInquiry(context.Background(), db.PostInquiryParams{ Usrid: req.Usrid, Title: req.Title, Body: req.Body }); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"response": res,
			})
		}
	})
	r.POST("/inquiries/reply/", func(c *gin.Context) {
		var req ReplyInquiryRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		if res, err := queries.ReplyInquiry(context.Background(), db.ReplyInquiryParams{ ID: req.ID, Reply: sql.NullString{ String: req.Reply, Valid: true } }); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"response": res,
			})
		}
	})

	r.GET("/items/get/", func(c *gin.Context) {
		if items, err := queries.GetItems(context.Background()); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"items": items,
			})
		}
	})
	r.POST("/items/post/", func(c *gin.Context) {
		var req PostItemRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		if res, err := queries.PostItem(context.Background(), req.Name); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"response": res,
			})
		}
	})

	r.POST("/items/delete", func(c *gin.Context) {
		var req struct {
			ID int32 `json:"id" binding:"required"`
		}
	
		// リクエストボディからIDをバインド
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid request body",
				"error":   err.Error(),
			})
			return
		}
	
		// アイテム削除クエリを実行
		err := queries.DeleteItem(context.Background(), req.ID)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to delete item",
				"error":   err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Item deleted successfully",
			})
		}
	})
	
	
	r.GET("/users-items/get/", func(c *gin.Context) {
		if users_items, err := queries.GetUsersItems(context.Background()); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"users_items": users_items,
			})
		}
	})
	r.POST("/users-items/post-to/", func(c *gin.Context) {
		var req PostItemToUserRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		if res, err := queries.PostItemToUser(context.Background(), db.PostItemToUserParams{ Usrid: req.Usrid, Itmid: req.Itmid, Amount: req.Amount }); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"response": res,
			})
		}
	})
	r.POST("/users-items/post-all/", func(c *gin.Context) {
		var req PostItemToAllUsersRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		if res, err := queries.PostItemToAllUsers(context.Background(), db.PostItemToAllUsersParams{ Itmid: req.Itmid, Amount: req.Amount }); err != nil {
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
