package dao

import (
	"github.com/dmzlingyin/utils/database/mongo"
	"github.com/dmzlingyin/utils/ioc"
)

func init() {
	ioc.Put(mongo.NewDatabase, "database")
	ioc.Put(mongo.NewScope, "scope")
	ioc.Put(NewHelloDao, "dao.hello")
}
