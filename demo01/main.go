package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type BlogUser struct {
	gorm.Model
	Username    string  `gorm:"size:50;not null"`
	Age         uint8   `gorm:"not null;default:20"`
	Price       float32 `gorm:"not null;precision:5;scale:2"`
	Description string  `gorm:"size:255;not null"`
	GroupID     uint    `gorm:"not null;comment:组ID"`
}

type BlogGroup struct {
	gorm.Model
	GroupName string `gorm:"size:20;not null;default:小学组"`
}

func main() {
	// 与数据建立连接
	dsn := "root:123456@tcp(106.13.223.242:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	var users []BlogUser
	// 子查询
	type S struct {
		AvgAge float32
	}
	subQuery := db.Table("blog_users").Select("avg(age) as avg_age")
	db.Where("age > (?)", subQuery).Find(&users)
	fmt.Printf("%+v", users)
	// 原生SQL
	//db.Raw("select * from blog_users").Find(&users)

}
