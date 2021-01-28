package logic

import (
	"context"

	"STFrontground-backend/rpc/tools/internal/svc"
	"STFrontground-backend/rpc/tools/tools"

	"github.com/tal-tech/go-zero/core/logx"
)

type StoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoryLogic {
	return &StoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StoryLogic) Story(in *tools.GetStoryReq) (*tools.GetStoryResp, error) {
	story, err := l.svcCtx.StoriesModel.FindOneRandom()
	if err != nil {
		return nil, err
	}
	return &tools.GetStoryResp{
		Story:  story.Story,
		Author: story.Auth,
	}, nil
}
