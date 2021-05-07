package core

import (
	"backend/internal/pkg/config"
	"backend/pkg/cache"
	"backend/pkg/db"
	"backend/pkg/ginx"
	"backend/pkg/logx"
	"backend/pkg/token"
)

func NewApp(vp *config.Viper) *ginx.App {
	return &ginx.App{
		Addr:  vp.Cfg.GetString("APP.addr"),
		Pprof: vp.Cfg.GetString("APP.pprof"),
		Mode:  vp.Cfg.GetString("APP.dev"),
	}
}

func NewMySQLConfig(vp *config.Viper) *db.MySQLConfig {
	return &db.MySQLConfig{
		Host:     vp.Cfg.GetString("DB.host"),
		Port:     vp.Cfg.GetString("DB.port"),
		User:     vp.Cfg.GetString("DB.user"),
		Password: vp.Cfg.GetString("DB.password"),
		DbName:   vp.Cfg.GetString("DB.DbName"),
	}
}

func NewRedisConfig(vp *config.Viper) *cache.RedisConfig {
	return &cache.RedisConfig{
		Addr:     vp.Cfg.GetString("REDIS.addr"),
		Password: vp.Cfg.GetString("REDIS.password"),
		DB:       vp.Cfg.GetInt("REDIS.db"),
	}
}

func NewLogConfig(vp *config.Viper) *logx.LogConfig {
	return &logx.LogConfig{
		Level:       vp.Cfg.GetString("LOG.level"),
		Filepath:    vp.Cfg.GetString("LOG.filepath"),
		FileMaxSize: vp.Cfg.GetInt("LOG.filemaxsize"),
		FileMaxAge:  vp.Cfg.GetInt("LOG.filemaxage"),
		MaxBackups:  vp.Cfg.GetInt("LOG.maxbackups"),
		Compress:    vp.Cfg.GetBool("LOG.compress"),
		Kfc: &logx.KafKaConfig{
			Addr:        vp.Cfg.GetStringSlice("KAFKA.addr"),
			Enable:      vp.Cfg.GetBool("KAFKA.enable"),
			Topics:      vp.Cfg.GetString("KAFKA.topics"),
			RequiredAck: vp.Cfg.GetInt("KAFKA.requiredack"),
			Mode:        vp.Cfg.GetString("KAFKA.mode"),
		},
	}
}

func NewJWTConfig(vp *config.Viper) *token.JsonWebTokenConfig {
	return &token.JsonWebTokenConfig{
		ExpireTime: vp.Cfg.GetInt64("JWT.expireTime"),
		Secret:     vp.Cfg.GetString("JWT.secret"),
		Audience:   vp.Cfg.GetString("JWT.audience"),
	}
}