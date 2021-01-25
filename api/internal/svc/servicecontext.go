package svc

import (
	"STFrontground-backend/api/internal/config"
	"STFrontground-backend/rpc/tools/toolsclient"
	"STFrontground-backend/rpc/user/userclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	User   userclient.User
	Tools  toolsclient.Tools
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userclient.NewUser(zrpc.MustNewClient(c.User)),
		Tools:  toolsclient.NewTools(zrpc.MustNewClient(c.Tools)),
	}
}
