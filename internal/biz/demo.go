package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Demo struct {
	Hello string
}

type DemoRepo interface {
	Save(context.Context, *Demo) (*Demo, error)
	Update(context.Context, *Demo) (*Demo, error)
	FindByID(context.Context, int64) (*Demo, error)
	ListByHello(context.Context, string) ([]*Demo, error)
	ListAll(context.Context) ([]*Demo, error)
}

type DemoUsecase struct {
	repo DemoRepo
	log  *log.Helper
}

func NewDemoUsecase(repo DemoRepo, logger log.Logger) *DemoUsecase {
	return &DemoUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *DemoUsecase) CreateGreeter(ctx context.Context, g *Demo) (*Demo, error) {
	uc.log.WithContext(ctx).Infof("CreateDemo: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
