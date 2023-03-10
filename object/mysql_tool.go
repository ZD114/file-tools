package object

import (
	"xorm.io/xorm"
	"zhangda/file-tools/config"
)

var _dbs map[string]*xorm.Engine

func init() {
	_dbs = make(map[string]*xorm.Engine, len(config.GetConfig().Spring.Datasource))

	for k, v := range config.GetConfig().Spring.Datasource {
		if engine, err := xorm.NewEngine("mysql", v.Url); err != nil {
			panic("连接数据库失败, error=" + err.Error())
		} else {
			engine.SetMaxOpenConns(v.MaxOpenConn)
			engine.SetMaxIdleConns(v.MaxIdleConn)

			_dbs[k] = engine
		}
	}
}

func GetCodePacketDB() *xorm.Engine {
	return _dbs["code-packet"]
}
