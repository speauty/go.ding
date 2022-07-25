package go_ding

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Verify struct {
	VerifyKey VerifyType
	VerifyVal string
}

type DRClient struct {
	channel    string
	webhook    string
	drType     DRType
	msgType    MsgType
	Verify     *Verify
	msgContent interface{}
}

func (client *DRClient) Ding() error {
	req, err := client.buildPostRequest()
	if err != nil {
		return err
	}
	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("DRClient::Ding client do err:%s", err)
	}
	return client.resp(res)
}

func (client *DRClient) setAttrChannel(channel string) {
	client.channel = channel
}

func (client *DRClient) setAttrWebHook(hook string) {
	client.webhook = hook
}

func (client *DRClient) setAttrDRType(drType DRType) {
	client.drType = drType
}

func (client *DRClient) setAttrMsgType(msgType MsgType) {
	client.msgType = msgType
}

func (client *DRClient) setAttrMsgContent(content interface{}) {
	client.msgContent = content
}

func (client *DRClient) setAttrVerify(verify *Verify) {
	client.Verify = verify
}

func (client *DRClient) buildPostRequest() (*http.Request, error) {
	body, _ := json.Marshal(client.msgContent)
	queryStr := ""
	if client.Verify.VerifyKey&VTKeySign == 1 {
		timestamp := time.Now().Unix()
		str2Sign := fmt.Sprintf("%d\\n%s", timestamp, client.Verify.VerifyVal)
		h := hmac.New(sha256.New, []byte("rr"))
		h.Write([]byte(str2Sign))
		sign := url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
		queryStr = fmt.Sprintf("&timestamp=%d&sign=%s", timestamp, sign)
	}
	req, err := http.NewRequest("POST", client.webhook+queryStr, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("DRClient::buildRequest err:%s", err)
	}
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func (client *DRClient) resp(resp *http.Response) error {
	defer func() { _ = resp.Body.Close() }()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("DRClient::resp ioutil.ReadAll err:%s", err)
	}

	return client.dingResp(content)
}

func (client *DRClient) dingResp(content []byte) error {
	resStruct := &DRErr{}

	if err := json.Unmarshal(content, resStruct); err != nil {
		return fmt.Errorf("DRClient::dingResp json.Unmarshal err:%s", err)
	}
	if resStruct.Code != 0 {
		return errors.New(resStruct.Msg)
	}
	return nil
}
