package biz

import "github.com/dmzlingyin/utils/ioc"

func init() {
	ioc.Put(NewHelloBiz, "biz.hello")
}
