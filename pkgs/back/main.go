package main

import (
	"context"
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"server/db"
	"strconv"
)

type UpdateStaminaRequest struct {
	ID      int32 `json:"id"      binding:"required"`
	Stamina int32 `json:"stamina" binding:"required"`
}

type UpdateConsecutiveDaysRequest struct {
	ID              int32 `json:"id"              binding:"required"`
	ConsecutiveDays int32 `json:"consecutiveDays" binding:"required"`
}

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

// 最終ログイン更新のリクエスト型
type UpdateLastLoginRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func main() {
	conn, err := sql.Open("postgres", "host=gpfdb port=5432 user=postgres password=password dbname=db sslmode=disable")
	if err != nil {
		panic(err)
	}
	queries := db.New(conn)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8080",
		},
		AllowMethods: []string{
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
		if res, err := queries.PostAnnounce(context.Background(), db.PostAnnounceParams{Title: req.Title, Body: req.Body}); err != nil {
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
	r.GET("/users/get-id-by-name/", func(c *gin.Context) {
		// クエリパラメータからユーザー名を取得
		userName := c.Query("name")

		if userName == "" {
			c.JSON(400, gin.H{
				"message": "User name is required",
			})
			return
		}

		// ユーザー名によるID取得クエリを実行
		userId, err := queries.GetUserIdByName(context.Background(), userName)
		if err != nil {
			if err == sql.ErrNoRows {
				// ユーザーが見つからない場合
				c.JSON(404, gin.H{
					"message": "User not found",
				})
			} else {
				// その他のデータベースエラー
				c.JSON(500, gin.H{
					"message": err.Error(),
				})
			}
			return
		}

		// ユーザーIDを返す
		c.JSON(200, gin.H{
			"userId": userId,
		})
	})

	r.POST("/users/post/", func(c *gin.Context) {
		var req PostUserRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		// ユーザーを作成
		res, err := queries.PostUser(context.Background(), req.Name)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		// user_detail に初期情報を挿入
		if err := queries.InsertUserDetail(context.Background(), res.ID); err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to create user_detail",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"response": res,
		})
	})

	r.POST("/users/ban-or-unban/", func(c *gin.Context) {
		var req BanOrUnbanUserRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		if res, err := queries.BanOrUnbanUser(context.Background(), db.BanOrUnbanUserParams{ID: req.ID, Banned: *req.Banned}); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"response": res,
			})
		}
	})

	// main() 関数内に新しいルートを追加
	r.POST("/users/update-last-login/", func(c *gin.Context) {
		var req UpdateLastLoginRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		if res, err := queries.UpdateUserLastLogin(context.Background(), req.ID); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"response": res,
			})
		}
	})

	r.GET("/users/get-last-login/", func(c *gin.Context) {
		// クエリパラメータからユーザーIDを取得
		usridStr := c.Query("usrid")
		usrid, err := strconv.ParseInt(usridStr, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "無効なユーザーID",
				"error":   err.Error(),
			})
			return
		}

		// ユーザーの最終ログイン日時を取得
		lastLogin, err := queries.GetUserLastLogin(context.Background(), int32(usrid))
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(404, gin.H{
					"message": "ユーザーが見つかりません",
				})
			} else {
				c.JSON(500, gin.H{
					"message": err.Error(),
				})
			}
			return
		}

		c.JSON(200, gin.H{
			"last_login": lastLogin,
		})
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
		if res, err := queries.PostInquiry(context.Background(), db.PostInquiryParams{Usrid: req.Usrid, Title: req.Title, Body: req.Body}); err != nil {
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
		if res, err := queries.ReplyInquiry(context.Background(), db.ReplyInquiryParams{ID: req.ID, Reply: sql.NullString{String: req.Reply, Valid: true}}); err != nil {
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

	r.GET("/users-items/get-by-user/", func(c *gin.Context) {
		// クエリパラメータからユーザーIDを取得
		usridStr := c.Query("usrid")
		usrid, err := strconv.ParseInt(usridStr, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid user ID",
				"error":   err.Error(),
			})
			return
		}

		// ユーザー別アイテム取得クエリを実行
		if users_items, err := queries.GetItemsByUser(context.Background(), int32(usrid)); err != nil {
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
		if res, err := queries.PostItemToUser(context.Background(), db.PostItemToUserParams{Usrid: req.Usrid, Itmid: req.Itmid, Amount: req.Amount}); err != nil {
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
		if res, err := queries.PostItemToAllUsers(context.Background(), db.PostItemToAllUsersParams{Itmid: req.Itmid, Amount: req.Amount}); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"response": res,
			})
		}
	})
	// スタミナ取得
	r.GET("/users/stamina/", func(c *gin.Context) {
		userIDStr := c.Query("id")
		userID, err := strconv.ParseInt(userIDStr, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"message": "Invalid user ID", "error": err.Error()})
			return
		}

		stamina, err := queries.GetUserStamina(context.Background(), int32(userID))
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
		} else {
			c.JSON(200, gin.H{"stamina": stamina})
		}
	})

	// 連続日数取得
	r.GET("/users/consecutive-days/", func(c *gin.Context) {
		userIDStr := c.Query("id")
		userID, err := strconv.ParseInt(userIDStr, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"message": "Invalid user ID", "error": err.Error()})
			return
		}

		consecutiveDays, err := queries.GetUserConsecutiveDays(context.Background(), int32(userID))
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
		} else {
			c.JSON(200, gin.H{"consecutive_days": consecutiveDays})
		}
	})

	// スタミナ更新
	r.POST("/users/update-stamina/", func(c *gin.Context) {
		var req UpdateStaminaRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		res, err := queries.UpdateUserStamina(context.Background(), db.UpdateUserStaminaParams{
			ID:      req.ID,
			Stamina: req.Stamina,
		})
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
		} else {
			c.JSON(200, gin.H{"response": res})
		}
	})

	// 連続日数更新
	r.POST("/users/update-consecutive-days/", func(c *gin.Context) {
		var req UpdateConsecutiveDaysRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		res, err := queries.UpdateUserConsecutiveDays(context.Background(), db.UpdateUserConsecutiveDaysParams{
			ID:              req.ID,
			ConsecutiveDays: req.ConsecutiveDays,
		})
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
		} else {
			c.JSON(200, gin.H{"response": res})
		}
	})

	r.Run(":63245")
}
