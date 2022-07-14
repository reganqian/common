package utils

import (
	"github.com/astaxie/beego/config"
	"time"
	"net/url"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha1"
	"strings"
	"errors"
	"gopkg.in/gomail.v2"
	"strconv"
	"github.com/sirupsen/logrus"
	"crypto/tls"
)
var (
	email_host = ""
	email_port = ""
	email_user = ""
	email_pwd = ""
	email_key = ""
	email_secret = ""
	email_name = ""
	email_region = ""
	smtp_host = ""
	
)
var Log = logrus.New()
func init() {
	//  = regan.qian@aliyun.com
	//  = money123!@#
	//  = C://project/src/authgrpc/util/

	iniconf, err1 := config.NewConfig("ini", "conf/config.ini")
	if err1 != nil {
		Log.Error(err1.Error())
	}


	// 2. 通过对象获取数据
	email_host = iniconf.String("email::email_host")
	smtp_host = iniconf.String("email::smtp_host")
	email_port = iniconf.String("email::email_port")
	email_user = iniconf.String("email::email_user")
	email_pwd = iniconf.String("email::email_pwd")
	email_key = iniconf.String("email::email_key")
	email_secret = iniconf.String("email::email_secret")
	email_name = iniconf.String("email::email_name")
	email_region = iniconf.String("email::email_region")
	
}

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}
func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}


//根据模板发送邮件
func SendEmailByHttp(mailTo, subject, emailCode string) (bool, error) {
	r := NewRequest([]string{mailTo}, subject, emailCode)
	body := emailCode
	r.body = body
	err := r.SendEmail()
	if err != nil {
		Log.Info(err)
		return false, err
	} else {
		return true, nil
	}
}

func (r *Request) SendEmailByHttp() error {
	err := SendToMail(email_user, email_host, r.subject, r.body, r.to)
    return err
}

func SendToMail(user, host, subject, body string, to []string) error {
	apiUrl := "http://"+host+"/?"


	// timeData := time.Now().UTC().UTC().Format("2006-01-02T15:04:05-0700")
	timeData := time.Now().UTC().UTC().Format(time.RFC3339)
	Log.Info(timeData)
	signatureNonce := GetIdStr()
	signData := "AccessKeyId=" + email_key
	signData += "&AccountName="+ user
	signData += "&Action=SingleSendMail"
	signData += "&AddressType=1"
	signData += "&Format=json"

	signData += "&HtmlBody=" + body
	signData += "&RegionId=" +email_region
	signData += "&ReplyToAddress=true"
	signData += "&SignatureMethod=HMAC-SHA1"
	signData += "&SignatureNonce=" + signatureNonce
	signData += "&SignatureVersion=1.0"
	signData += "&Subject=" + subject
	signData += "&Timestamp=" + timeData
	signData += "&ToAddress=" + to[0]
	signData += "&Version=2017-06-22"
	signData += "&FromAlias=" + email_name

	v := url.Values{}
	v.Add("AccessKeyId",email_key)
	v.Add("AccountName", user)
	v.Add("Action","SingleSendMail")
	v.Add("AddressType","1")
	v.Add("Format","json")
	v.Add("HtmlBody", body)
	v.Add("RegionId",email_region)
	v.Add("ReplyToAddress","true")
	v.Add("SignatureMethod","HMAC-SHA1")
	v.Add("SignatureNonce", signatureNonce)
	v.Add("SignatureVersion","1.0")
	v.Add("Subject",subject)
	v.Add("Timestamp", timeData)
	v.Add("ToAddress", to[0])
	v.Add("Version","2017-06-22")
	v.Add("FromAlias", email_name)
	encData := v.Encode()
	
	encData = strings.Replace(encData, "%2B", "%20", -1)
	encData = strings.Replace(encData, "%", "%25", -1)
	encData = strings.Replace(encData, "+", "%20", -1)
	encData = strings.Replace(encData, "*", "%2A", -1)
	encData = strings.Replace(encData, "=", "%3D", -1)
	encData = strings.Replace(encData, "%7E", "~", -1)
	encData = strings.Replace(encData, "&", "%26", -1)

	apiUrl += signData
	apiUrl += "&Signature=" + getSignature("POST&%2F&" + encData)
	
	resData, err := DoGet(apiUrl)
	if resData["EnvId"] != nil {
		return nil
	} else {
		if resData["Code"] == nil {
			err = errors.New("send email failed")
		} else {
			err = errors.New(resData["Code"].(string))
		}
		
	}

	return err
}

func getSignature(signData string) string {
	key := []byte(email_secret + "&")
	// signData = base64.StdEncoding.EncodeToString([]byte(signData))
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(signData))
	signedStr := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return signedStr
}



//根据模板发送邮件
func SendEmailByContent(mailTo, subject, emailBody string) (bool, error) {
	r := NewRequest([]string{mailTo}, subject, emailBody)

	err := r.SendEmail()
	if err != nil {
		Log.Info(err)
		return false, err
	} else {
		return true, nil
	}
}

func (r *Request) SendEmail() error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
    mailConn := map[string]string {
		"user":  email_user, 
        "pass":  email_pwd,  
        "host":  smtp_host,
        "port": email_port,
    }

    port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

    m := gomail.NewMessage()
	// m.SetHeader("From","YHXD" + "<" + mailConn["user"] + ">")  //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("From",mailConn["user"])
    m.SetHeader("To", r.to...)  //发送给多个用户
    m.SetHeader("Subject", r.subject)  //设置邮件主题
    m.SetBody("text/html", r.body)     //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    err := d.DialAndSend(m)
    return err
}

