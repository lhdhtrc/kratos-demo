package data

import (
	"context"

	"demo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

func NewDemoRepo(data *Data, logger log.Logger) biz.DemoRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Demo) (*biz.Demo, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Demo) (*biz.Demo, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Demo, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Demo, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Demo, error) {
	return nil, nil
}
