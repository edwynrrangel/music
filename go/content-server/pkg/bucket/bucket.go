package bucket

import (
	"context"
	"io"
)

type Strategy interface {
	DownloadFile(ctx context.Context, bucketName, fileName string) (io.ReadCloser, error)
}

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) DownloadFile(ctx context.Context, bucketName, fileName string) (io.ReadCloser, error) {
	return c.strategy.DownloadFile(ctx, bucketName, fileName)
}
