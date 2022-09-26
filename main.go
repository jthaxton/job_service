package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
		r := gin.Default()
		db,err := InitializeDb()
		handler := Handler{Db: db}
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("HERE")
		r.POST("/create", handler.AddJobToQueue)
		r.GET("/next", handler.GetNextJob)
		r.Run()
}
