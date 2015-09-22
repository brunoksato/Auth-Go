package db

func init(){
    db, _ := gorm.Open("postgres", "dbname=auth sslmode=disable")
    db.DB()
    return db;
}



