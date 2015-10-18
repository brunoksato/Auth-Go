package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "github.com/gin-gonic/gin"
    _ "fmt"
    "net/http"
    _ "strings"
)

var (
    db            gorm.DB
    sqlConnection string
    err           error
)

type User struct {
    Name     string    `json:"name"`
    Email    string    `json:"email"`
    Login    string    `json:"login"`
    Password string    `json:"password"`
}

func Authentication(c *gin.Context){
    var json User
    if c.BindJSON(&json) == nil {
        result := db.Where("Login = ? and Password >= ?", json.Login, json.Password).Find(&json)
        if result.RowsAffected > 0 {
            c.JSON(http.StatusOK, gin.H{"status": result})
        } else {
            c.JSON(http.StatusNotFound, gin.H{"status": result})
        }
    }
}

func Register(c *gin.Context){
    var json User
    if c.BindJSON(&json) == nil {
        result := db.Where("Login = ? and Password >= ?", json.Login, json.Password).Find(&json)
        if result.RowsAffected == 0 {
            db.NewRecord(json)
            db.Create(&json)
            c.JSON(http.StatusOK, gin.H{"status": "Ok"})
        } else{
            c.JSON(http.StatusNotFound, gin.H{"status": "Error"})
        }

    }
}

func UsersList(c *gin.Context){
    users := []User{}
    result := db.Find(&users)
    c.JSON(http.StatusOK, gin.H{"status": result})
}

func connection() {
    sqlConnection = "dbname=auth sslmode=disable"

    db, err = gorm.Open("postgres", sqlConnection)
    db.DB()
    db.SingularTable(true)

    if err != nil {
        panic(err)
        return
    }
}

func main() {
    //create connection db
    connection()

    router := gin.Default()
    router.Static("/", "./public")

    router.POST("/auth/login", Authentication)
    router.POST("/api/register", Register)
    router.POST("/api/users", UsersList)

    router.Run(":8080")
}
