//go:build wireinject

package main

import (
	"github.com/asynccnu/be-infoSum/grpc"
	"github.com/asynccnu/be-infoSum/ioc"
	"github.com/asynccnu/be-infoSum/pkg/grpcx"
	"github.com/asynccnu/be-infoSum/repository/cache"
	"github.com/asynccnu/be-infoSum/repository/dao"
	"github.com/asynccnu/be-infoSum/service"
	"github.com/google/wire"
)

func InitGRPCServer() grpcx.Server {
	wire.Build(
		ioc.InitGRPCxKratosServer,
		grpc.NewInfoSumServiceServer,
		service.NewInfoSumService,
		cache.NewRedisInfoSumCache,
		dao.NewMysqlInfoSumDAO,
		// 第三方
		ioc.InitDB,
		ioc.InitRedis,
		ioc.InitLogger,
		ioc.InitEtcdClient,
	)
	return grpcx.Server(nil)
}
