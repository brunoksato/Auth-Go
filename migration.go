package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
)

type Model struct {
    ID uint `gorm:"primary_key"`
}

type User struct {
    gorm.Model
    Name string
    Email string
    Login string
    Password string
}

func main(){

    db, _ := gorm.Open("postgres", "dbname=auth sslmode=disable")
    db.DB()
    db.SingularTable(true)

    db.AutoMigrate(&User{})
    db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
    db.AutoMigrate(&User{})
}