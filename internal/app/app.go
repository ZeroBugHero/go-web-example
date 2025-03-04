package app

import (
	"goWebExample/api/rest/handlers"
	"goWebExample/internal/middleware"
	"goWebExample/internal/pkg/db"
	myZap "goWebExample/internal/pkg/zap"
	"goWebExample/internal/repository/user"
	"goWebExample/internal/service/user_service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// NewGin 创建并配置一个新的 Gin 引擎实例
func NewGin(logger *zap.Logger) *gin.Engine {
	engine := gin.New()
	// 加载中间件
	middleware.LoadMiddleware(logger, engine)
	return engine
}

// 中间件相关依赖

// DatabaseSet 数据库相关依赖
var DatabaseSet = wire.NewSet(db.NewGormConfig)

// LoggerSet 日志相关依赖
var LoggerSet = wire.NewSet(
	myZap.NewZap,
)

// 业务模块相关依赖

var (
	// RepositorySet 仓储层依赖
	RepositorySet = wire.NewSet(
		user.NewUserRepository,
		// 其他仓储
	)

	// ServiceSet 服务层依赖
	ServiceSet = wire.NewSet(
		user_service.NewUserService,
		// 其他服务
	)

	// HandlerSet Handler层依赖
	HandlerSet = wire.NewSet(
		handlers.NewUserHandler,
		// 其他Handler
	)

	// ProviderSet 汇总所有依赖
	ProviderSet = wire.NewSet(
		RepositorySet,
		ServiceSet,
		HandlerSet,
	)
)
