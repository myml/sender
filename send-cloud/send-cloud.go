package sendcloud

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"bj.git.sndu.cn/deepinid/sender/share"
	"github.com/pkg/errors"
)

func NewSendCloud(account Account) *SendCloud {
	return &SendCloud{
		account: account,
		client:  http.DefaultClient,
		server:  SendCloudServer,
	}
}

type Account struct {
	SMS_USER string
	SMS_KEY  string
}
type SendCloud struct {
	account Account
	client  *http.Client
	server  string
}

func (sc *SendCloud) SendSMS(opt *share.SendSMSOption) error {
	resp, err := sc.client.PostForm(sc.server+"/smsapi/send", sc.toParam(opt))
	if err != nil {
		return errors.Wrap(err, "http")
	}
	defer resp.Body.Close()
	err = sc.parserInfo(resp.Body, nil)
	if err != nil {
		return errors.Wrap(err, "parser")
	}
	return nil
}

func (sc *SendCloud) parserInfo(r io.Reader, info interface{}) error {
	type Result struct {
		Result     bool        `json:"result"`
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message"`
		Info       interface{} `json:"info"`
	}
	var result Result
	result.Info = info
	err := json.NewDecoder(r).Decode(&result)
	if err != nil {
		errors.Wrap(err, "decode")
	}
	if result.StatusCode != 200 {
		return errors.Errorf("%s(%d)", result.Message, result.StatusCode)
	}
	return nil
}

func (sc *SendCloud) toParam(opt *share.SendSMSOption) url.Values {
	param := url.Values{}
	param.Set("smsUser", sc.account.SMS_USER)
	param.Set("templateId", opt.Template)

	phones := make([]string, 0, len(opt.Phone))
	for i := range opt.Phone {
		phone := opt.Phone[i]
		phoneLength := len(phone)
		if phoneLength < share.MinPhoneLength || phoneLength > share.MaxPhoneLength {
			continue
		}
		if phone[0] == '+' {
			if strings.HasPrefix(phone, "+86") {
				// delete china area code
				phone = phone[3:]
			} else {
				// set message type to International SMS
				param.Set("msgType", fmt.Sprint(msgTypeInternationalSMS))
			}
		}
		phones = append(phones, phone)
	}
	param.Set("phone", strings.Join(phones, ","))

	if opt.TemplateVariable != nil {
		b, err := json.Marshal(opt.TemplateVariable)
		if err != nil {
			panic(err)
		}
		param.Set("vars", string(b))
	}
	param.Set("signature", sc.signature(param))
	return param
}

func (sc *SendCloud) signature(param url.Values) string {
	keys := make([]string, 0, len(param))
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	data := make([]string, 0, len(keys)+2)
	data = append(data, sc.account.SMS_KEY)
	for i := range keys {
		data = append(data, fmt.Sprintf("%s=%s", keys[i], param.Get(keys[i])))
	}
	data = append(data, sc.account.SMS_KEY)

	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(data, "&"))))
}
