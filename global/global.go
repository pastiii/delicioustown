package global

import (
	"DeliciousTown/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 常量
const (
	ConfigFile = "./config.yaml" // 配置文件
)

// 变量
var (
	GvaConfig      config.ServerConfig // 全局配置
	DefaultLogger  *zap.Logger
	UserLogger     *zap.Logger
	GvaMysqlClient *gorm.DB //Mysql客户端
	GvaRedis       *redis.Client
)
