package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"os"
)
type Test struct {
	//gorm.Model
	Name      string                `gorm:"uniqueIndex:udx_name;type:varchar(10)"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name"`
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3302)/lqhandsome?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(255)
	}
	db.AutoMigrate(&Test{})
	db.Debug().Create(&Test{Name: "123456789刘"})
	t := []Test{}
	db.Debug().Find(&t)
	fmt.Println(t)
}

func (t *Test)Preloading(tx *gorm.DB)error{

	tx = tx.Where("1 =1")
	return  nil
}


func (t *Test)BeforeSave(tx *gorm.DB)error{

	//t = new(Test)
	//t.Name =time.Now().Format("20060102150304")
	return  nil
}