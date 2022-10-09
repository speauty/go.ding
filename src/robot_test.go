package src

import (
	"fmt"
	"github.com/speauty/go.ding/src/cfg"
	"github.com/speauty/go.ding/src/consts"
	"github.com/speauty/go.ding/src/types"
	"testing"
)

const webhook = "https://oapi.dingtalk.com/robot/send?access_token=4248b970034950aaa018882ba6b49bf4fc153bfe925e7af05c8680e2f221c676"

func TestGoDRText(t *testing.T) {
	currentCfg := &cfg.Cfg{}
	currentCfg.Default()
	currentCfg.Channel = "tt"
	drClient := &DRClient{}
	msg := types.NewDRMsgText()
	msg.SetContent(fmt.Sprintf("[%s]%s", currentCfg.Channel, "hello"))

	drClient.setAttrWebHook(webhook)
	drClient.setAttrChannel(currentCfg.Channel)
	drClient.setAttrDRType(currentCfg.DRType)
	drClient.setAttrMsgType(msg.GetMsgType())
	drClient.setAttrVerify(&Verify{VerifyKey: consts.VTKey, VerifyVal: "tt"})
	drClient.setAttrMsgContent(msg)

	err := drClient.Ding()
	if err != nil {
		t.Error(err)
	}
}

func TestGoDRMD(t *testing.T) {
	currentCfg := &cfg.Cfg{}
	currentCfg.Default()
	currentCfg.Channel = "tt"
	drClient := &DRClient{}
	msg := types.NewDrMsgMarkdown()
	msg.SetTitle("测试")
	msg.SetContent("#### 杭州天气 @150XXXXXXXX \n> 9度，西北风1级，空气良89，相对温度73%\n> ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n> ###### 10点20分发布 [天气](https://www.dingalk.com) \n")
	drClient.setAttrWebHook(webhook)
	drClient.setAttrChannel(currentCfg.Channel)
	drClient.setAttrDRType(currentCfg.DRType)
	drClient.setAttrMsgType(msg.GetMsgType())
	drClient.setAttrVerify(&Verify{VerifyKey: consts.VTKey, VerifyVal: "tt"})
	drClient.setAttrMsgContent(msg)

	err := drClient.Ding()
	if err != nil {
		t.Error(err)
	}
}

func TestGoDRLink(t *testing.T) {
	currentCfg := &cfg.Cfg{}
	currentCfg.Default()
	currentCfg.Channel = "tt"
	drClient := &DRClient{}
	msg := types.NewDRMsgLink()
	msg.SetTitle(fmt.Sprintf("[%s]%s", currentCfg.Channel, "测试-链接消息"))
	msg.SetContent("吧啦啦啦啊巴拉克列发附件或第三方节好地方")
	msg.SetPicUrl("https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png")
	msg.SetMessageUrl("https://open.dingtalk.com/document/")

	drClient.setAttrWebHook(webhook)
	drClient.setAttrChannel(currentCfg.Channel)
	drClient.setAttrDRType(currentCfg.DRType)
	drClient.setAttrMsgType(msg.GetMsgType())
	drClient.setAttrVerify(&Verify{VerifyKey: consts.VTKey, VerifyVal: "tt"})
	drClient.setAttrMsgContent(msg)

	err := drClient.Ding()
	if err != nil {
		t.Error(err)
	}
}

func TestGoDRActionCard(t *testing.T) {
	currentCfg := &cfg.Cfg{}
	currentCfg.Default()
	currentCfg.Channel = "tt"
	drClient := &DRClient{}
	msg := types.NewDRMsgActionCard()
	msg.SetTitle("测试-卡片消息")
	msg.SetContent("![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png) \n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划")
	msg.SetSingleTitle("阅读全文")
	msg.SetSingleURL("https://www.dingtalk.com/")

	drClient.setAttrWebHook(webhook)
	drClient.setAttrChannel(currentCfg.Channel)
	drClient.setAttrDRType(currentCfg.DRType)
	drClient.setAttrMsgType(msg.GetMsgType())
	drClient.setAttrVerify(&Verify{VerifyKey: consts.VTKey, VerifyVal: "tt"})
	drClient.setAttrMsgContent(msg)

	err := drClient.Ding()
	if err != nil {
		t.Error(err)
	}
}
