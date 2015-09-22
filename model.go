package main

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