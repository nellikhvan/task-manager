package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nellikhvan/task-manager/models"
)

var db *gorm.DB

func main() {
	// Initialize the database
	initDB()

	// Initialize the Gin router
	r := gin.Default()

	// Set up routes
	r.GET("/tasks", getTasks)
	r.GET("/tasks/:id", getTask)
	r.POST("/tasks", createTask)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	// Start the server
	r.Run(":8080")
}

func initDB() {
	// Open a database connection
	var err error
	db, err = gorm.Open("postgres", "host=postgres port=5432 user=user dbname=task_manager password=password sslmode=disable")
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Task{})
}


func getTasks(c *gin.Context) {
	var tasks []models.Task
	db.Find(&tasks)
	c.JSON(200, tasks)
}

func getTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, task)
	}
}

func createTask(c *gin.Context) {
	var task models.Task
	c.BindJSON(&task)

	db.Create(&task)
	c.JSON(200, task)
}

func updateTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&task)
	db.Save(&task)
	c.JSON(200, task)
}

func deleteTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task
	d := db.Where("id = ?", id).Delete(&task)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}