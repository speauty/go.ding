package cfg

import "github.com/speauty/go.ding/src/consts"

type Cfg struct {
	Channel string        // 渠道
	DRType  consts.DRType // 默认自定义机器人
}

func (c *Cfg) Default() {
	c.Channel = "dev"
	c.DRType = consts.CustomRobot
}
