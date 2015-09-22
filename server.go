package main

import (
    _ "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    //"github.com/gocraft/web"
    "github.com/gin-gonic/gin"
    //"github.com/corneldamian/json-binding"
    _ "fmt"
    "net/http"
    _ "strings"
)

type User struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Login string `json:"login"`
    Password string `json:"password"`
}

// func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
//     c.User =
//     next(rw, req)
// }

// func (c *Context) Authentication(rw web.ResponseWriter, req *web.Request) {
//     //user := ctx.RequestJSON.(*Authenticate)
//     //fmt.Println("User " + c.User.Name);
//     // ctx.ResponseJSON = binding.SuccessResponse("User " + c.User.Name)
//     //ctx.ResponseStatus = http.StatusUnauthorized
// }

// // func UsersList(ctx *Context, rw web.ResponseWriter, req *web.Request) {

// // }

// func UserCreate(ctx *Context, rw web.ResponseWriter, req *web.Request) {
//     user := ctx.RequestJSON.(*User)
//     ctx.ResponseJSON = binding.SuccessResponse("User " + user.Name)
// }
//

func Authentication(c *gin.Context){
    var json User
    if c.BindJSON(&json) == nil {
        c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
        // if json.User == "manu" && json.Password == "123" {
        //     c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
        // } else {
        //     c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        // }
    }
}

func Register(c *gin.Context){
    var json User
    if c.BindJSON(&json) == nil {
        c.JSON(http.StatusOK, gin.H{"status": json.Name})
        // if json.User == "manu" && json.Password == "123" {
        //     c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
        // } else {
        //     c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        // }
    }
}


func main() {

    router := gin.Default()
    router.Static("/", "./public")

    router.POST("/auth/login", Authentication)
    router.POST("/api/register", Register)
    router.GET("/api/users", UsersList)

    router.Run(":8080")
    //router := web.New(Context{})

    //static
//    router.Middleware(web.StaticMiddleware("public"))

    //
    // router.Middleware(binding.Response(nil))
    // router.Middleware(binding.Request(Authenticate{}, nil)).Post("/auth/login", Authentication)

    // router.Middleware(binding.Response(nil))
    // router.Middleware(binding.Request(User{}, nil)).Post("/api/users", UserCreate)

    //user api
    //router.Post("/auth/login", (*Context).Authentication)
    //router.Post("/api/users", (*Context).UserCreate)
    // router.Put("/users/:id", (*Context).UsersUpdate)
    // router.Delete("/users/:id", (*Context).UsersDelete)

   // http.ListenAndServe("localhost:3000", router)
}