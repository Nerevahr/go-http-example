package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
	"net/http"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"log"
)

type Module struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Courses []Course
}

type Course struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

func main() {
	username := "postgres"
	password := "password"
	host := "localhost"
	dbname := "test"
	
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, dbname, password))
	
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}
	defer db.Close()
	
    db.AutoMigrate(&Module{}, &Course{})

	r := gin.Default()
	r.POST("/modules", createModule(db))	
    r.GET("/modules", getModules(db))
    r.Run()
}

func getModules(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var modules []Module
        db.Preload("Courses").Find(&modules)
        c.JSON(200, modules)
    }
}

func createModule(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var json Module
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        db.Create(&json)
        c.JSON(http.StatusOK, gin.H{"success": json})
    }
}
