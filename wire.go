//go:build wireinject

package main

import (
	"github.com/asynccnu/be-infosum/grpc"
	"github.com/asynccnu/be-infosum/ioc"
	"github.com/asynccnu/be-infosum/pkg/grpcx"
	"github.com/asynccnu/be-infosum/repository/cache"
	"github.com/asynccnu/be-infosum/repository/dao"
	"github.com/asynccnu/be-infosum/service"
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
