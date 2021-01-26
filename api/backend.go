package main

import (
	"STFrontground-backend/api/internal/config"
	"STFrontground-backend/api/internal/handler"
	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/rpc/pkg/errcode"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

// 返回的结构体，json格式的body
type Message struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

// 定义错误处理函数
func errorHandler(e error) (int, interface{}) {
	err := e.(*errcode.Error)
	return http.StatusConflict, Message{
		Code: err.Code(),
		Desc: err.Msg(),
	}
}

var configFile = flag.String("f", "etc/backend-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(rest.RestConf{
		Port: c.Port,
		ServiceConf: service.ServiceConf{
			Log:  logx.LogConf{Path: c.LogConfig.LogPath}, // 日志路径
			Mode: c.LogConfig.LogMode,
		},
	})
	defer server.Stop()

	// 设置错误处理函数
	httpx.SetErrorHandler(errorHandler)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
