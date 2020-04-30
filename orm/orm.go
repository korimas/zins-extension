package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris/v12"
)

type Orm struct {
	app *iris.Application
	DB  *gorm.DB
}

func New(app *iris.Application) *Orm {
	orm := Orm{
		app: app,
	}
	orm.Connect()
	return &orm
}

func (orm *Orm) Connect() {
	db, err := gorm.Open("sqlite3", "/tmp/test.db")
	if err != nil {
		fmt.Printf("errutils: %s", err)
		panic("连接数据库失败 %s")
	}
	fmt.Printf("Connect DB success")
	orm.DB = db
}
