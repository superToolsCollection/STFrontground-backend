package logic

import (
	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
)

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
	client := &http.Client{}
	url := l.svcCtx.Config.VerificationCode.URL
	info := make(map[string]string)
	info["apikey"] = l.svcCtx.Config.VerificationCode.APIKey
	info["mobile"] = req.Mobile
	info["text"] = fmt.Sprintf(l.svcCtx.Config.VerificationCode.Text, 1234)

	bytesData, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json;charset=utf-8")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d",
				resp.StatusCode)
	}
	result, err := ioutil.ReadAll(ioutil.NopCloser(resp.Body))
	return &types.GetVerificationCodeResp{
		ResultCode: string(result),
	}, nil
}
