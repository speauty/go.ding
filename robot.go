package go_ding

type DRClient struct {
	channel    string
	drType     DRType
	msgType    MsgType
	msgContext string
}

func (client *DRClient) SetAttrChannel(channel string) {
	client.channel = channel
}

func (client *DRClient) SetAttrDRType(drType DRType) {
	client.drType = drType
}

func (client *DRClient) SetAttrMsgType(msgType MsgType) {
	client.msgType = msgType
}

func (client *DRClient) SetAttrMsgContext(msg string) {
	client.msgContext = msg
}

func (client *DRClient) GetAttrChannel() string {
	return client.channel
}

func (client *DRClient) GetAttrDRType() DRType {
	return client.drType
}

func (client *DRClient) GetAttrMsgType() MsgType {
	return client.msgType
}

func (client *DRClient) GetAttrMsgContext() string {
	return client.msgContext
}
