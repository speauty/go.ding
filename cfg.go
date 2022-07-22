package go_ding

type Cfg struct {
	Channel string // 渠道
	DRType  DRType // 默认自定义机器人
}

func (c *Cfg) Default() {
	c.Channel = "dev"
	c.DRType = CustomRobot
}
