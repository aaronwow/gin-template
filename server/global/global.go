package global

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"gandi.icu/marketing-form/server/config"
)

var (
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	GVA_CONFIG config.Server

	BlackCache local_cache.Cache
)
