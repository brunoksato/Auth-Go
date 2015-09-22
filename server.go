package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "github.com/gocraft/web"
    "github.com/corneldamian/json-binding"
    _ "fmt"
    "net/http"
    _ "strings"
)

type Context struct {
    RequestJSON interface{}
    ResponseJSON interface{}
    ResponseStatus int
}

type Authenticate struct {
    Username string
    Password string
}

func Authentication(ctx *Context, rw web.ResponseWriter, req *web.Request) {
    user := ctx.RequestJSON.(*Authenticate)
    ctx.ResponseJSON = binding.SuccessResponse("User " + user.Username + " Pass: " +  user.Password)
    //ctx.ResponseStatus = http.StatusUnauthorized
}

func UsersList(ctx *Context, rw web.ResponseWriter, req *web.Request) {

}

func UsersCreate(ctx *Context, rw web.ResponseWriter, req *web.Request) {

}


func main() {

    db, err := gorm.Open("postgres", "user=brunosato dbname=auth sslmode=disable")
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
    db.SingularTable(true)

    router := web.New(Context{})

    //static
    router.Middleware(web.StaticMiddleware("public"))

    //
    router.Middleware(binding.Response(nil))
    router.Middleware(binding.Request(Authenticate{}, nil)).Post("/auth/login", Authentication)

    //user api
    // router.Get("/users", (*Context).UsersList)
    // router.Post("/users", (*Context).UsersCreate)
    // router.Put("/users/:id", (*Context).UsersUpdate)
    // router.Delete("/users/:id", (*Context).UsersDelete)

    http.ListenAndServe("localhost:3000", router)
}