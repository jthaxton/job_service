package main

import (
	"database/sql"
	"encoding/json"
	"log"
	// "fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Handler struct {
	Db  Database
}
type Database struct {
	Conn *sql.DB
}

type Body struct {
	CustomId string `json:"custom_id"`
	Kind     string `json:"kind"`
	DataJson interface{}  `json:"data_json"`
}

// Body{
// 	custom_id string
// 	kind string
// 	data_json json
// }
func InitializeDb() (Database, error) {
	db := Database{}
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 		"localhost", 5432, "root", "example", "jobs_db")
	url := "postgresql://root:example@postgres/job_db?sslmode=disable"
	conn, err := sql.Open("postgres", url)
	if err != nil {
			log.Println(err.Error())
			return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
			log.Println(err.Error())
			return db, err
	}
	log.Println("Database connection established")
	return db, nil
}

func (handler *Handler) AddJobToQueue(ctx *gin.Context) {
	body := Body{}
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(403,map[string]string{"error": err.Error()})
		return
	}
	var id int
	var createdAt string
	query := `INSERT INTO jobs (custom_id, kind, data_json) VALUES ($1, $2, $3) RETURNING id, created_at`
	// str := fmt.Sprintf("%v",body.DataJson)
	// raw := json.RawMessage(str)
	b, err := json.Marshal(body.DataJson)
	if err != nil {
		ctx.JSON(403,map[string]string{"error": err.Error()})
		return
	}
	err = handler.Db.Conn.QueryRow(query, body.CustomId, body.Kind, b).Scan(&id, &createdAt)
	if err != nil {
			ctx.JSON(403,map[string]string{"error": err.Error()})
			return
	}

	ctx.JSON(200, map[string]int{"id": id})
}

func (handler *Handler) GetNextJob(ctx *gin.Context) {
	body := Body{}
	ctx.ShouldBind(body)
	job := Job{}
	query := `SELECT * FROM jobs LIMIT 1`
	err := handler.Db.Conn.QueryRow(query).Scan(&job.ID, &job.CreatedAt, &job.CustomId, &job.Kind, &job.DataJson)
	if err != nil {
			ctx.JSON(403,map[string]string{"error": err.Error()})
			return
	}

	ctx.JSON(200, map[string]Job{"job": job})
}