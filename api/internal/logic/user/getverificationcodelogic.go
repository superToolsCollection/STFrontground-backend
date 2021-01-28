package logic

import (
	"STFrontground-backend/api/internal/util"
	"STFrontground-backend/rpc/pkg/errcode"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type YunPianResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (y YunPianResp) String() string {
	return fmt.Sprintf("code:%d\n"+
		"msg:%s\n",
		y.Code, y.Msg)
}

type GetVerificationCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetVerificationCodeLogic {
	return GetVerificationCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVerificationCodeLogic) GetVerificationCode(req types.GetVerificationCodeReq) (*types.GetVerificationCodeResp, error) {
	u := l.svcCtx.Config.VerificationCode.URL
	vCode := util.GenerateVerificationCode()
	data := url.Values{
		"apikey": {l.svcCtx.Config.VerificationCode.APIKey},
		"mobile": {req.Mobile},
		"text":   {fmt.Sprintf(l.svcCtx.Config.VerificationCode.Text, vCode)}}
	request, err := http.PostForm(u, data)
	if err != nil {
		return nil, errcode.NewHttpRequestError
	}
	defer request.Body.Close()

	request.Header.Add("Accept", "application/json;charset=utf-8")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	result, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, errcode.ReadRequestError
	}
	if request.StatusCode != http.StatusOK {
		return nil, errcode.HttpBadRequestError
	}
	yunpianResp := &YunPianResp{}
	err = json.Unmarshal(result, yunpianResp)
	if err != nil {
		return nil, errcode.JsonUnMarshalError
	}
	return &types.GetVerificationCodeResp{
		ResultCode: yunpianResp.Code,
		Msg:        yunpianResp.Msg,
		VCode:      vCode,
	}, nil
}
