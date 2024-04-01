package biz

type HelloBiz interface {
	Hello() string
}

func NewHelloBiz() HelloBiz {
	return &helloBiz{}
}

type helloBiz struct {
}

func (b *helloBiz) Hello() string {
	return "hello"
}
