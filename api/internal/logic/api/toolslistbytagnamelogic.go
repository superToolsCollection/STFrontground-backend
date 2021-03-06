package logic

import (
	"STFrontground-backend/rpc/tools/tools"
	"context"

	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ToolsListByTagNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewToolsListByTagNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) ToolsListByTagNameLogic {
	return ToolsListByTagNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ToolsListByTagNameLogic) ToolsListByTagName(req types.ToolsListByTagNameReq) (*types.ToolsListByTagNameResp, error) {
	result, err := l.svcCtx.Tools.GetToolListByTagName(l.ctx, &tools.GetToolListByTagNameReq{
		PageSize: req.PageSize,
		Page:     req.Page,
		TagName:  req.TagName,
	})
	if err != nil {
		return nil, err
	}

	t := make([]*types.Tool, len(result.Tools))
	for i := 0; i < len(result.Tools); i++ {
		t[i].ToolId = result.Tools[i].ToolID
		t[i].Name = result.Tools[i].Name
		t[i].Api = result.Tools[i].Api
		t[i].ApiDescribe = result.Tools[i].ApiDescribe
		t[i].Picture = result.Tools[i].Picture
		tags := make([]*types.Tag, len(result.Tools[i].Tags))
		for j := 0; j < len(result.Tools[i].Tags); j++ {
			t[i].Tags[j].TagId = result.Tools[i].Tags[j].TagID
			t[i].Tags[j].TagName = result.Tools[i].Tags[j].TagName
		}
		t[i].Tags = tags
	}
	return &types.ToolsListByTagNameResp{
		Tools: t,
	}, nil
}
