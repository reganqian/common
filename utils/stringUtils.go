package utils

import (
	"math/rand"
	"time"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/base64"
	"github.com/rs/xid"
	"strconv"
	"fmt"
	"strings"
	// . "yunapi/log"
)

func UniqAppend(dataList []string, data string) ([]string) {
	for _, data1 := range dataList {
		if data1 ==  data {
			return dataList
		} 
	}
	dataList = append(dataList, data)
	return dataList
}

func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

var r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func  GetRandomNum(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < l; i++ {
		br := r.Intn(10)
		result = append(result, bytes[br])
	}
	return string(result)
}


func CreateMd5String(addStr, salt string) string {

	m5 := md5.New()

	m5.Write([]byte(addStr))
	m5.Write([]byte(salt))

	st := m5.Sum(nil)	
	token := hex.EncodeToString(st)
	
	return string(token)
}

func Md5String(addStr string) string {
	m5 := md5.New()
	m5.Write([]byte(addStr))
	st := m5.Sum(nil)	
	token := hex.EncodeToString(st)
	return string(token)
}


func ProdPwd(oldPwd, salt string) string {
	oldKey := Md5String(salt)
	key := []byte(strings.ToTitle(oldKey[0 : 8]))
	result, err := DesEncrypt([]byte(oldPwd), key)
	if err != nil {
		panic(err)
	}
	hexstr := fmt.Sprintf("%X", result)
	fmt.Println(hexstr)
	return hexstr
}



// Convert json string to map
func JsonToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		
		return nil, err
	}
	
	return m, nil
}
 
// Convert map json string
func MapToJson(m map[string]string) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {

		return "", nil
	}
 
	return string(jsonByte), nil
}


func MakeTimestamp() int64 {

    return time.Now().UTC().UnixNano() / int64(time.Millisecond)
}

//生成订单号：YYYYMMDDHHmmssSSS + gameId + pp + 
func CreateOrderNo(orderPre string) string {
	date := GetTodyDay()
	data := GetTimeTick64()
	code := fmt.Sprintf("%s%s%s%s", orderPre, date, data, GetRandomNum(3))
	return code
}


func GetTodyDay() string {
	timeTemplate := "060102"
	return time.Now().UTC().Format(timeTemplate)
}


func GetTimeTick64() string {
	intData := time.Now().UnixNano() / 1e6
	s  := strconv.FormatInt(intData, 10)
	content := s[4 : len(s)-1]
	return content
}


func GetIdStr() string {
	return xid.New().String()
}

func FormatTimeToStr(t time.Time) string {
	timeTemplate1 := "2006-01-02 15:04:05"
	return  t.Format(timeTemplate1)
}

func GetTodaySandE(t time.Time) (s, e int64) {
	
	timeStr := t.Format("2006-01-02")	
	loc := time.FixedZone("UTC-8", 8 * 60 * 60)
	st, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	s = st.Unix()
	e = s + 86399
	
	return s, e
}

func GetBeiJingTime(t time.Time) time.Time {
	loc := time.FixedZone("UTC-8", 8 * 60 * 60)
	t = t.In(loc)
	
	return t
}

func TimeIntToTimeStr(timeInt int64) string {
	loc := time.FixedZone("UTC-8", 8 * 60 * 60)
	d := time.Unix(timeInt, 0).In(loc)
	timeTemplate := "2006-01-02 15:04:05"
	return d.Format(timeTemplate)

}

func TimeIntToStr(timeInt int64) string {
	loc := time.FixedZone("UTC-8", 8 * 60 * 60)
	d := time.Unix(timeInt, 0).In(loc)
	return GetTimeStr(d)
}

func TimeIntToMontStr(timeInt int64) string {
	loc := time.FixedZone("UTC-8", 8 * 60 * 60)
	d := time.Unix(timeInt, 0).In(loc)
	return GetTimeMonthStr(d)
}

func GetLocalTime() time.Time {
	loc := time.FixedZone("UTC-8", 8 * 60 * 60)
	nowData := time.Now().In(loc)
	return nowData
}

func FormatTimeToStrWithZone(t time.Time, zone int) string {
	timeTemplate1 := "2006-01-02 15:04:05"
	loc := time.FixedZone("UTC-8", zone * 60 * 60)
	t = t.In(loc)
	return  t.Format(timeTemplate1)
}

func GetTimeUnix(t time.Time) int64 {
	return t.UTC().Unix()
}


func GetTodayStr() string {
	timeTemplate := "20060102"
	return time.Now().UTC().Format(timeTemplate)
}

func GetTimeStr(inTime time.Time) string {
	timeTemplate := "2006-01-02"
	return inTime.Format(timeTemplate)
}

func GetTimeMonthStr(inTime time.Time) string {
	timeTemplate := "2006-01"
	return inTime.Format(timeTemplate)
}

func GetNowStr() string {
	timeTemplate := "20060102150405"
	return time.Now().UTC().Format(timeTemplate)
}



func GetProxyTag(userId uint32, ipId string) string {
	userIdStr := strconv.Itoa(int(userId))
	str := "ext.at(\"group\") == \""+userIdStr+"\" &&  ext.at(\"ipId\")==\""+ipId+"\"" 
	return GetBase64Data(str)
}

func GetProxyGroup(userId uint32) string {
	userIdStr := strconv.Itoa(int(userId))
	str := "ext.at(\"group\") == \""+userIdStr+"\"" 
	return GetBase64Data(str)
}

func GetProxyGroupNot64(userId uint32) string {
	userIdStr := strconv.Itoa(int(userId))
	str := "ext.at(\"group\") == \""+userIdStr+"\"" 
	return str
}

func GetBase64Data(data string) string {
	input := []byte(data)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return string(encodeString)
}

func GetTodayZeroTime() time.Time {
	timeTemplate1 := "2006-01-02 15:04:05"
	timeStr := GetLocalTime().Format("2006-01-02")
	t2 := timeStr + " 00:00:00"
	stamp, _ := time.ParseInLocation(timeTemplate1, t2, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return stamp
}

func FormatStringToInt32(data string) (int32, error) {
	j,err := strconv.ParseInt(data,10,32)
	if err != nil {
		return 0, err
	}
	return int32(j), nil
}

//获取一天前的字符串
func GetTimeBefore(inTime time.Time) time.Time {
	next := inTime.Add(- time.Hour * 24)
	return next
}

// func GetZeroTime(inTime time.Time) int64 {
// 	return time.Date(inTime.Year(), inTime.Month(), inTime.Day(), 0, 0, 0, 0, inTime.Location()).Unix()
// }

func GetTimeBeforeMonth(inTime time.Time) time.Time {
	next := inTime.AddDate(0, -1, 0)
	return next
}

//strconv.Itoa(int(i))
func FormatInt32ToStr(data int32) (string) {
	return strconv.Itoa(int(data))
}


//是否是不重复的数组
func CheckSameFriend(dataList []string) bool {
	fmap := make(map[string]string)
	for _, fid := range dataList {
		dbD := fmap[fid]
		if dbD == "" {
			fmap[fid] = fid
		} else {
			return false
		}
	}
	return true
} 

func CheckInStrings(data string, datas []string) bool {
	for _, v := range datas {
		if v == data {
			return true
		}
	}

	return false
}



func StringsToUint32s(accidList []string) []uint32 {
	var intList []uint32
	for _, accStr := range accidList {
		data, _ := strconv.Atoi(accStr)
		intList = append(intList, uint32(data))
	}
	return intList
}

func StringsToUints(accidList []string) []uint {
	var intList []uint
	for _, accStr := range accidList {
		data, _ := strconv.Atoi(accStr)
		intList = append(intList, uint(data))
	}
	return intList
}

func StrToTime(str string) time.Time {

	var timeLayoutStr = "2006-01-02 15:04:05"
	st, _ := time.Parse(timeLayoutStr, str) //string转time

   	return st
}


func Int64ToTime(timeInt int64) time.Time {

	return time.Unix(timeInt, 0)
}

func RecerseList(x []interface{}) ([]interface{}) {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
        x[i], x[j] = x[j], x[i]
    }
	return x
}


//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day() + 1)
	return GetZeroTime(d)
}
//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}
 
//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func GetLastMonthStartEnd(t time.Time) (time.Time, int64, int64) {
	// now := time.Now()
	lastMonthFirstDay := t.AddDate(0, -1, - t.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, t.Location()).Unix()
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, t.Location()).Unix()
	return lastMonthFirstDay, lastMonthStart, lastMonthEnd
 }

func HidePrivateInfo(data string) string {
	if len(data) < 2 {
		return data
	} else if len(data) < 6 {
		return string(data[0:2]) + "****"
	} else {
		return string(data[0:3]) + "***" + string(data[len(data)-3:])
	}
}


//获取在as 不在bs的数据
func GetDiffUint32s(as []uint32, bs []uint32) (cs []uint32) {
	isIn := false
	for _, a := range as {
		for _, b := range bs {
			if a == b {
				isIn = true
			}
		}
		if !isIn {
			cs = append(cs, a)
		}
	}
	return cs
}

func Uint32ToString(data uint32) string {
	dataStr := strconv.Itoa(int(data))
	return dataStr
}
