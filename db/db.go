package db

import (
	"github.com/muxiyun/MAE/models"
	"github.com/go-xorm/core"
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/xorm"
	"github.com/prometheus/common/log"
)

var DBEngine xorm.Engine

func init(){
	//连接数据库
	DBEngine,err:=xorm.NewEngine("sqlite3","./test.db")
	if err!=nil{
		log.Fatal("connect database failed!")
	}

	//设置数据库相关名称与结构相关名称映射规则
	DBEngine.SetMapper(core.SameMapper{})

	//同步数据库
	DBEngine.Sync2(new(models.User))
}