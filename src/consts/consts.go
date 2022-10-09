package consts

// DRType 钉钉机器人类型(DingRobot)
type DRType uint8

const (
	// InteriorRobot 企业内部机器人
	InteriorRobot DRType = 1 << iota
	// GroupTplRobot 群模板机器人
	GroupTplRobot
	// CustomRobot 自定义机器人
	CustomRobot
)

// MsgType 消息类型
type MsgType uint8

const (
	MsgText MsgType = 1 << iota
	MsgLink
	MsgMarkdown
	MsgActionCard
	MsgFeedCard
	MsgEmpty
)

// VerifyType 验证类型
type VerifyType uint8

const (
	VTKey VerifyType = 1 << iota
	VTKeySign
)

const (
	BtnOrientationCol uint8 = 0 // 按钮竖直排列
	BtnOrientationRow uint8 = 1 // 按钮横向排列
)
