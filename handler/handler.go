package handler

import "github.com/dmzlingyin/utils/ioc"

func init() {
	ioc.Put(NewHelloHandler, "handler.hello")
}
