package main

import (
	"context"
	"database/sql"

	"server/db"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := sql.Open("postgres", "host=gpfdb port=5432 user=postgres password=password dbname=db sslmode=disable")
	if err != nil {
		panic(err)
	}
	queries := db.New(conn)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		if players, err := queries.ListPlayers(context.Background()); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"players": players,
			})
		}
	})
	r.Run(":63245")
}
