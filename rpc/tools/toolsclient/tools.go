// Code generated by goctl. DO NOT EDIT!
// Source: tools.proto

//go:generate mockgen -destination ./tools_mock.go -package toolsclient -source $GOFILE

package toolsclient

import (
	"context"

	"STFrontground-backend/rpc/tools/tools"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	MorseReq                 = tools.MorseReq
	GetStoryResp             = tools.GetStoryResp
	Tool                     = tools.Tool
	GetToolListReq           = tools.GetToolListReq
	GetToolListByTagNameReq  = tools.GetToolListByTagNameReq
	GetToolListByTagNameResp = tools.GetToolListByTagNameResp
	GetStoryReq              = tools.GetStoryReq
	MorseResp                = tools.MorseResp
	QrCodeReq                = tools.QrCodeReq
	QrCodeResp               = tools.QrCodeResp
	Rgb2HexResp              = tools.Rgb2HexResp
	Tag                      = tools.Tag
	GetToolListResp          = tools.GetToolListResp
	Rgb2HexReq               = tools.Rgb2HexReq

	Tools interface {
		GetToolList(ctx context.Context, in *GetToolListReq) (*GetToolListResp, error)
		GetToolListByTagName(ctx context.Context, in *GetToolListByTagNameReq) (*GetToolListByTagNameResp, error)
		Morse(ctx context.Context, in *MorseReq) (*MorseResp, error)
		QrCode(ctx context.Context, in *QrCodeReq) (*QrCodeResp, error)
		Rgb2Hex(ctx context.Context, in *Rgb2HexReq) (*Rgb2HexResp, error)
		Story(ctx context.Context, in *GetStoryReq) (*GetStoryResp, error)
	}

	defaultTools struct {
		cli zrpc.Client
	}
)

func NewTools(cli zrpc.Client) Tools {
	return &defaultTools{
		cli: cli,
	}
}

func (m *defaultTools) GetToolList(ctx context.Context, in *GetToolListReq) (*GetToolListResp, error) {
	client := tools.NewToolsClient(m.cli.Conn())
	return client.GetToolList(ctx, in)
}

func (m *defaultTools) GetToolListByTagName(ctx context.Context, in *GetToolListByTagNameReq) (*GetToolListByTagNameResp, error) {
	client := tools.NewToolsClient(m.cli.Conn())
	return client.GetToolListByTagName(ctx, in)
}

func (m *defaultTools) Morse(ctx context.Context, in *MorseReq) (*MorseResp, error) {
	client := tools.NewToolsClient(m.cli.Conn())
	return client.Morse(ctx, in)
}

func (m *defaultTools) QrCode(ctx context.Context, in *QrCodeReq) (*QrCodeResp, error) {
	client := tools.NewToolsClient(m.cli.Conn())
	return client.QrCode(ctx, in)
}

func (m *defaultTools) Rgb2Hex(ctx context.Context, in *Rgb2HexReq) (*Rgb2HexResp, error) {
	client := tools.NewToolsClient(m.cli.Conn())
	return client.Rgb2Hex(ctx, in)
}

func (m *defaultTools) Story(ctx context.Context, in *GetStoryReq) (*GetStoryResp, error) {
	client := tools.NewToolsClient(m.cli.Conn())
	return client.Story(ctx, in)
}
