package main

import (
	"gandi.icu/marketing-form/server/core"
	"gandi.icu/marketing-form/server/global"
	"gandi.icu/marketing-form/server/initialize"
)

func main() {
	// initialize config
	global.GVA_VP = core.Viper()
	// initialize JWT expires time & buffer time
	initialize.OtherInit()
	// initialize zap logger
	global.GVA_LOG = core.Zap()
}
