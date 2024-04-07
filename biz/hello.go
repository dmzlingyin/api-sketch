package biz

import (
	"api-sketch/dao"
	"context"
)

type HelloBiz interface {
	Hello(ctx context.Context, name string) error
}

func NewHelloBiz(hd dao.HelloDao) HelloBiz {
	return &helloBiz{
		hd: hd,
	}
}

type helloBiz struct {
	hd dao.HelloDao
}

func (b *helloBiz) Hello(ctx context.Context, name string) error {
	return b.hd.Create(ctx, name)
}
