package biz

import (
	"api-sketch/dao"
	"context"
)

type HelloBiz interface {
	Hello(ctx context.Context, name string) error
	Ping(ctx context.Context, name string) (string, error)
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

func (b *helloBiz) Ping(ctx context.Context, name string) (string, error) {
	res, err := b.hd.Query(ctx, name)
	if err != nil {
		return "", err
	}
	if res != nil {
		return res.ID.Hex(), nil
	}
	return "", nil
}
