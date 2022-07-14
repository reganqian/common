package utils

import (
	"strings"
	"fmt"
	"errors"
)

func GetSecretInfo(salt string) (string, string) {
	secretId := GetIdStr()
	secretKey := ProdPwd(secretId, salt)
	return secretId, secretKey
}

func CreateSign(orderId, totalNumStr, currency, timeStr, secretId, secretKey  string) string {
	preStr := orderId + "|" + totalNumStr + "|" + currency + "|" + timeStr + "|" + secretId + "|" + secretKey
	md5Str := Md5String(preStr)
	return strings.ToLower(md5Str)
}

func CheckSign(flowId, orderId, totalNumStr, currency, timeStr, secretId, secretKey string) string {
	preStr := flowId + "|" + orderId + "|" + totalNumStr + "|" + currency + "|" + timeStr + "|" + secretId + "|" + secretKey
	md5Str := Md5String(preStr)
	return strings.ToLower(md5Str)
}

func NotifySign(detail, secretKey string) string {
	md5Str := ProdPwd(detail, secretKey)
	sign := strings.ToLower(md5Str)
	data := detail + "#" + sign
	return data
}

func CheckNotifySign(detail, secretKey string) (string, error) {
	details := strings.Split(detail, "#")
	if len(details) != 2 {
		fmt.Println("收到通知, 数据格式异常", detail)
		return "", errors.New("数据错误")
	}
	info := details[0]
	sign := details[1]
	md5Str := ProdPwd(info, secretKey)
	thisSign := strings.ToLower(md5Str)
	if sign != thisSign {
		fmt.Println("收到通知, 签名错误", detail)
		return "", errors.New("签名错误")
	}

	return info, nil
} 

func GetQuerySign(queryTime, secretId, secretKey string) string{
	preStr := secretId + "|" + queryTime
	md5Str := ProdPwd(preStr, secretKey)
	return strings.ToLower(md5Str)
}

