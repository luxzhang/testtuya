package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"log"
	"testtuya/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

var GlobalConfigSetting = &Data{}

// Data .
type Data struct {
	db *sql.DB
}

// NewData .
func NewData(c *conf.Data) (d *Data, err error) {
	log.Println("-------init newdata------")
	//使用NewMySQL方法进行连接池对象的初始化
	d = new(Data)
	GlobalConfigSetting.db, _ = sql.Open(c.Database.Driver, c.Database.Source)
	log.Println("-------init newdata complete------")
	return &Data{}, nil
}