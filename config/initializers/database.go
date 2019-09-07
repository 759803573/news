package initializers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	//DB manage db
	DB = &dbConnecton{}
)

type dbConnecton struct {
	Conn *gorm.DB
}

//DBBaseModel base model
type DBBaseModel gorm.Model

func (d *dbConnecton) Init(typ string, kwargs map[string]string) (err error) {
	fmt.Println("test")
	d.Conn, err = gorm.Open("sqlite3", "test.db")
	return
}

func (d *dbConnecton) Close() {
	d.Conn.Close()
}
