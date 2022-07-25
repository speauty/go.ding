package go_ding

import (
	"fmt"
	"testing"
)

func TestGoDRText(t *testing.T) {
	currentCfg := &Cfg{}
	currentCfg.Default()
	currentCfg.Channel = "tt"
	drClient := &DRClient{}
	drClient.setAttrWebHook("https://oapi.dingtalk.com/robot/send?access_token=4248b970034950aaa018882ba6b49bf4fc153bfe925e7af05c8680e2f221c676")
	drClient.setAttrChannel(currentCfg.Channel)
	drClient.setAttrDRType(currentCfg.DRType)
	drClient.setAttrMsgType(MsgText)
	drClient.setAttrVerify(&Verify{
		VerifyKey: VTKey,
		VerifyVal: "tt",
	})
	text := &DRMsgText{MsgType: "text"}
	text.Text.Content = fmt.Sprintf("[%s]%s", currentCfg.Channel, "hello")
	drClient.setAttrMsgContent(text)
	err := drClient.Ding()
	t.Error(err)
}
