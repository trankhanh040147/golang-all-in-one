package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func main() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// DB_CONNECTION from launch.json
	// dsn := os.Getenv("DB_CONNECTION")
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3309)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(string(dsn)), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("DB Connection: ", db)

	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "Belajar Golang",
		Description: "Belajar Golang untuk membuat API",
		Status:      "Active",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	jsData, err := json.Marshal(item)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(jsData))

	jsString := `{"id":1,"title":"Belajar Golang","description":"Belajar Golang untuk membuat API","status":"Active","created_at":"2021-10-13T15:04:05Z","updated_at":"2021-10-13T15:04:05Z"}`

	var item2 TodoItem

	if err := json.Unmarshal([]byte(jsString), &item2); err != nil {
		log.Fatalln(err)
	}

	log.Println(item2)

	// -----------------------------------------------

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		items := v1.Group("items")
		{
			items.POST("")
			items.GET("")
			items.GET("/:id")
			items.PATCH("/:id")
			items.DELETE("/:id")
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3000") // Fix: Add a comma between `addr` and `":3000"`

	// fmt.Println("Hello, World!")
	// fmt.Println(os.Getenv("APP_NAME"))
}
